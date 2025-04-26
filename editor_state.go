package main

type EditorState struct {
	selectionIndex int
	editorData     map[int]EditorAsset
}
