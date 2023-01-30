package pdf

import "github.com/unidoc/unipdf/v3/common/license"

func Initialize() {
	err := license.SetMeteredKey("cc75cf1caa32c9e694404fbbe18d2a08c16e94378e0421d2f5b6bfc29be95c2b")
	if err != nil {
		panic(err)
	}
}
