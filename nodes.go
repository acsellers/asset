package asset

import (
	"text/template/parse"
)

func cacheNode(title string) *parse.ListNode {
	tn := &parse.TextNode{
		NodeType: parse.NodeText,
	}
	cachedAssetList[title] = tn

	return &parse.ListNode{
		NodeType: parse.NodeList,
		Nodes:    []parse.Node{tn},
	}
}

func checkCacheNode(title string) *parse.PipeNode {
	return &parse.PipeNode{
		NodeType: parse.NodePipe,
		Cmds: []*parse.CommandNode{
			&parse.CommandNode{
				NodeType: parse.NodeCommand,
				Args: []parse.Node{
					&parse.IdentifierNode{
						NodeType: parse.NodeIdentifier,
						Ident:    "assetCheckCache",
					},
					&parse.StringNode{
						NodeType: parse.NodeString,
						Quoted:   title,
						Text:     title,
					},
				},
			},
		},
	}
}

func newTextNode(s string) *parse.TextNode {
	return &parse.TextNode{
		NodeType: parse.NodeText,
		Text:     []byte(s),
	}
}

func actionNodeFor(ct string, r *require) *parse.ActionNode {
	switch r.Type {
	case "tree":
		return newTreeNode(ct, r.Item)
	case "directory":
		return newDirectoryNode(ct, r.Item)
	default:
		return newItemNode(ct, r.Item)
	}
}

func newTreeNode(ct, s string) *parse.ActionNode {
	funcName := "assetTree"
	stringNodes := []parse.Node{
		&parse.StringNode{
			NodeType: parse.NodeString,
			Quoted:   ct,
			Text:     ct,
		},
		&parse.StringNode{
			NodeType: parse.NodeString,
			Quoted:   s,
			Text:     s,
		},
	}
	switch ct {
	case "js":
		funcName = "assetJsTree"
		stringNodes = stringNodes[1:]
	case "css":
		funcName = "assetCssTree"
		stringNodes = stringNodes[1:]
	}
	nodes := []parse.Node{
		&parse.IdentifierNode{
			NodeType: parse.NodeIdentifier,
			Ident:    funcName,
		},
	}

	return &parse.ActionNode{
		NodeType: parse.NodeAction,
		Pipe: &parse.PipeNode{
			NodeType: parse.NodePipe,
			Cmds: []*parse.CommandNode{
				&parse.CommandNode{
					NodeType: parse.NodeCommand,
					Args:     append(nodes, stringNodes...),
				},
			},
		},
	}

}

func newDirectoryNode(ct, s string) *parse.ActionNode {
	return &parse.ActionNode{}
}

func newItemNode(ct, s string) *parse.ActionNode {
	funcName := "assetLoad"
	stringNodes := []parse.Node{
		&parse.StringNode{
			NodeType: parse.NodeString,
			Quoted:   ct,
			Text:     ct,
		},
		&parse.StringNode{
			NodeType: parse.NodeString,
			Quoted:   s,
			Text:     s,
		},
	}
	switch ct {
	case "js":
		funcName = "assetJsLoad"
		stringNodes = stringNodes[1:]
	case "css":
		funcName = "assetCssLoad"
		stringNodes = stringNodes[1:]
	}
	nodes := []parse.Node{
		&parse.IdentifierNode{
			NodeType: parse.NodeIdentifier,
			Ident:    funcName,
		},
	}

	return &parse.ActionNode{
		NodeType: parse.NodeAction,
		Pipe: &parse.PipeNode{
			NodeType: parse.NodePipe,
			Cmds: []*parse.CommandNode{
				&parse.CommandNode{
					NodeType: parse.NodeCommand,
					Args:     append(nodes, stringNodes...),
				},
			},
		},
	}
}

func newTree(name string, elseList *parse.ListNode) *parse.Tree {
	//TODO: add in checkCache branch
	return &parse.Tree{
		Root: elseList,
	}
}

func newListNode(nodes []parse.Node) *parse.ListNode {
	return &parse.ListNode{
		NodeType: parse.NodeList,
		Nodes:    nodes,
	}
}
