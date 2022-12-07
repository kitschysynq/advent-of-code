package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	tree := NewDirTree()

	p := NewParser(os.Stdin)
	if err := p.Parse(tree); err != nil {
		log.Fatalf("error parsing: %s", err.Error())
	}

	free := 70000000 - tree.Root.TotalSize()
	min := int(^uint(0) >> 1)

	if err := tree.Walk(func(n *Node) error {
		if !n.IsDir { return nil } // ignore files
		if s := n.TotalSize(); s + free >= 30000000 && s < min {
			min = s
		}
		return nil
	}); err != nil {
		log.Fatalf("error walking tree: %s", err.Error())
	}

	fmt.Println(min)
}

type Parser struct {
	s *bufio.Scanner
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: bufio.NewScanner(r)}
}

func (p *Parser) Parse(d *DirTree) error {
	for p.s.Scan() {
		line := p.s.Text()

		// ignore blank lines
		if len(line) == 0 {
			continue
		}

		for len(line) > 0 && line[0] == '$' {
			var err error
			line, err = p.processCommand(line[2:], d)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Parser) processCommand(c string, d *DirTree) (string, error) {
	cmd := c[0:2]
	switch cmd {
	case "cd":
		dir := c[3:]
		switch dir {
		case "..":
			d.Up()
			return "", nil
		default:
			return "", d.Cd(dir)
		}
	case "ls":
		for p.s.Scan() {
			next := p.s.Text()
			if next[0] == '$' {
				return next, nil
			}
			if err := p.processEntry(next, d); err != nil {
				return "", err
			}
		}
		return "", nil
	}

	return "", fmt.Errorf("unknown command %q", c)
}

func (p *Parser) processEntry(e string, d *DirTree) error {
	if strings.HasPrefix(e, "dir ") {
		return d.Mkdir(e[4:])
	}
	var size int
	var name string
	if n, err := fmt.Sscanf(e, "%d %s", &size, &name); err != nil || n != 2 {
		return err
	}
	return d.Creat(name, size)
}

type DirTree struct {
	Root *Node
	cur  *Node
}

func NewDirTree() *DirTree {
	return &DirTree{
		Root: &Node{
			Name:  "/",
			IsDir: true,
		},
	}
}

func (d *DirTree) Up() {
	if d.cur.parent != nil {
		d.cur = d.cur.parent
	}
}

func (d *DirTree) Cd(name string) error {
	if d.cur == nil {
		if d.Root.Name == name {
			d.cur = d.Root
			return nil
		}
		return fmt.Errorf("current dir has no subdir named %q", name)
	}
	for _, node := range d.cur.Children {
		if node.Name == name {
			if !node.IsDir {
				return fmt.Errorf("%q is not a directory", name)
			}
			d.cur = node
			return nil
		}
	}
	return fmt.Errorf("current dir has no subdir named %q", name)
}

func (d *DirTree) Mkdir(name string) error {
	if d.cur.hasChild(name) {
		return fmt.Errorf("%q already exists", name)
	}
	d.cur.Children = append(d.cur.Children, &Node{
		Name:   name,
		IsDir:  true,
		parent: d.cur,
	})
	return nil
}

func (d *DirTree) Creat(name string, size int) error {
	if d.cur.hasChild(name) {
		return fmt.Errorf("%q already exists", name)
	}
	d.cur.Children = append(d.cur.Children, &Node{
		Name:   name,
		Size:   size,
		parent: d.cur,
	})
	return nil
}

func (d *DirTree) Walk(w WalkFn) error {
	nodes := []*Node{d.Root}
	for len(nodes) > 0 {
		if err := w(nodes[0]); err != nil {
			return err
		}
		if len(nodes) == 1 {
			nodes = append([]*Node{}, nodes[0].Children...)
			continue
		}
		nodes = append(nodes[1:], nodes[0].Children...)
	}
	return nil
}

type WalkFn func(n *Node) error

type Node struct {
	Name     string
	Size     int
	IsDir    bool
	parent   *Node
	Children []*Node
}

func (n *Node) hasChild(name string) bool {
	if !n.IsDir {
		return false
	}
	for _, node := range n.Children {
		if node.Name == name {
			return true
		}
	}
	return false
}

func (n *Node) TotalSize() int {
	if n.IsDir {
		var s int
		for _, child := range n.Children {
			s += child.TotalSize()
		}
		return s
	}
	return n.Size
}
