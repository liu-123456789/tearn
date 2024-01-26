package main

v1, d1 := Silced([]int{1, 2, 3, 4, 5}, 2)
fmt.Printf("删除后的 %v,\n删除的下标 %d", v1, d1)

type Silcei interface {
	int
}

func Silcedeom[T Silcei](val []T, dex int) ([]T, int) {
	vala := []T{}
	if dex >= len(val) || dex < 0 {
		println("输入不合法")
	}
	for k, v := range val {
		if k != dex-1 {
			vala = append(vala, v)

		}

	}
	return vala, dex

}