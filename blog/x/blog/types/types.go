package types

// type Post struct {
// 	proto.Message
// 	Id        uint64
// 	Creator   string
// 	Title     string
// 	Body      string
// 	CreatedAt string
// 	Approved  bool
// }
// // имплементировали нашу структуру Post в интерфейс ProtoMarshal
// // для того чтобы соответсвовать методу Mustmarshal
// // func (p Post) Marshal() ([]byte, error) {
// // 	return nil, fmt.Errorf("Implement method Marshal this %s", "types.go, marshal function")
// // }

// // func (p Post) MarshalTo(data []byte) (n int, err error) {
// // 	return 0, fmt.Errorf("%s", "$")
// // }
// // func (p Post) MarshalToSizedBuffer(dAtA []byte) (int, error) {
// // 	return 0, fmt.Errorf("%s", "err$")
// // }
// // func (p Post) Size() int {
// // 	return 0
// // }
// // func (p Post) Unmarshal(data []byte) error {
// // 	return nil
// // }
