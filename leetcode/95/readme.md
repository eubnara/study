## golang composite literals

https://go.dev/ref/spec#Composite_literals


아래사항 모두 가능

```go
tmp := TreeNode{Val: start}
return []*TreeNode{
	&tmp,
}
```

```go
return []*TreeNode{
	&TreeNode{Val: start},
}

```

```go
return []*TreeNode{
	{Val: start},
}
```