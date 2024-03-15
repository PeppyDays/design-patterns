package prototype

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}
	folder2 := &Folder{
		children: []Inode{file2},
		name:     "Folder2",
	}

	folder1.print("  ")
	folder2.print("  ")

	clonedFolder := folder2.clone()
	clonedFolder.print("  ")
}
