package main

func main() {
	a := App{}
	a.Initialize()
	InitializeMDB()
	a.Run(":5000")
}
