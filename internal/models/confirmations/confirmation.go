package confirmations

import (
	"github.com/google/uuid"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Number string
type UID string
type Code string

func NewAuthCode(codeLen uint) Code {
	rand.Seed(time.Now().UnixNano())

	min := int64(math.Pow(10, float64(codeLen-1)))
	max := int64(math.Pow(10, float64(codeLen)) - 1)

	return Code(strconv.FormatInt(rand.Int63n(max-min+1)+min, 10))
}

func NewRequestUID() UID {
	return UID(uuid.New().String())
}

type Confirmation struct {
	RequestUID UID  `json:"requestId"`
	AuthCode   Code `json:"code"`
}

func NewConfirmation(codeLen uint) Confirmation {
	return Confirmation{
		RequestUID: NewRequestUID(),
		AuthCode:   NewAuthCode(codeLen),
	}
}

type CurrentUnixDate int64
