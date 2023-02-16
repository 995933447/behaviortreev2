package gobehaviortree

type (
	OnExecTreeFunc func(tree *Tree)
	OnTreeChildFinishFunc func(tree *Tree, result Result, root BaseNode)
)

//Tree tree describe
type Tree struct {
	root BaseNode
	onExec OnExecTreeFunc
	onChildFinish OnTreeChildFinishFunc
}

//NewTree new tree
func NewTree(onExec OnExecTreeFunc, onChildFinish OnTreeChildFinishFunc) *Tree {
	return &Tree{
		onExec: onExec,
		onChildFinish: onChildFinish,
	}
}

//SetRoot setroot
func (t *Tree) SetRoot(node BaseNode) {
	if t.root != nil {
		t.root.OnUninstall()
	}
	t.root = node
	t.root.SetTree(t)
	t.root.OnInstall()
}

//Run run
func (t *Tree) Run() {
	t.exec()
}

func (t *Tree) exec() {
	if t.onExec != nil {
		t.onExec(t)
	}

	res := t.root.OnEnter()
	t.root.OnExit()
	if t.onChildFinish != nil {
		t.onChildFinish(t, res, t.root)
	}
}

func (t *Tree) GetRoot() BaseNode {
	return t.root
}