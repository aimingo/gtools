# gtools
the go language development tools

### go get github.com/aimingo/gtools

##### gtools.SliceUNIQ Slice elements to remove duplicates

`gtools.SliceUNIQ([]uin64{1,1,2})`

remove duplicates result

`1,2`


#####  gtools.SliceExists ele is exists in the slice

`
gtools.SliceExists([]uin64{1,1,2},1)  => true
gtools.SliceExists([]uin64{1,1,2},3)  => false
`
