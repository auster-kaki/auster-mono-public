package austerid

import "github.com/rs/xid"

func Generate[T ~string]() T { return T(xid.New().String()) }
