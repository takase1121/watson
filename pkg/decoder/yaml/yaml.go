package yaml

import (
	"io"

	"gopkg.in/yaml.v2"

	"github.com/genkami/watson/pkg/decoder/util"
	"github.com/genkami/watson/pkg/vm"
)

func Decode(w io.Writer, val *vm.Value) error {
	obj := util.ToObject(val)
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	return enc.Encode(obj)
}
