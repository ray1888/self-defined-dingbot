package encoder

import (
	"testing"

	"github.com/ray1888/self-defined-dingbot/src/messages/middlemsg"
	"github.com/ray1888/self-defined-dingbot/src/parser/encode/impl"
)

type Test struct {
	TestName string
	Input    middlemsg.Body
	Output   string
}

func TestEncodeTextMsg(t *testing.T) {
	tests := []Test{
		{TestName: "Decode Push Event Success",
			Input: middlemsg.Body{
				EventType:    "push",
				Username:     "John Smith",
				Source:       "master",
				Project:      "Diaspora",
				CommitNumber: "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Commits: []middlemsg.Commit{
					{Number: "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
						Info: "Update Catalan translation to e38cb41."},
					{Number: "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
						Info: "fixed readme"},
				},
				TotalCommitCount: 2,
			},
			Output: "John Smith pushed to branch master at repository Diaspora\ncommit: da1560886d4f094c3e6c9ef40349f7d38b5d27d7 \ncommit-message: {{ .CommitMsg}} \n项目：Diaspora",
		},
	}
	for _, test := range tests {
		result, err := impl.EncodeTextMsg(test.Input)
		if err != nil {
			t.Fatalf("TestName %s encode msg meet error +%v", test.TestName, err)
		}
		if result != test.Output {
			t.Fatalf("TestName %s encode msg is not expected", test.TestName)
		}
	}

}

//func TestEncodeLinkMsg(t *testing.T) {
//	impl.EncodeLinkMsg()
//}
//
//func TestEncodeActionCardMsg(t *testing.T) {
//	impl.EncodeActionCardMsg()
//}