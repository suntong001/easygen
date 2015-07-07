package easygenapi

import (
	"github.com/danverbraganza/varcaser/varcaser"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// pre-configed varcaser caser converters
// the names are self-explanatory from their definitions
var (
	cls2lc = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.LowerCamelCase}
	cls2uc = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.UpperCamelCase}
	cls2ss = varcaser.Caser{From: varcaser.LowerSnakeCase, To: varcaser.ScreamingSnakeCase}

	ck2lc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerCamelCase}
	ck2uc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCase}
<<<<<<< HEAD
	ck2ls = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerSnakeCase}
=======
	ck2ls = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.lower_snake_case}
>>>>>>> 854eeb44cca3948e0c8fec92e675acd32545ea8b
	ck2ss = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.ScreamingSnakeCase}

	clc2ss = varcaser.Caser{From: varcaser.LowerCamelCase, To: varcaser.ScreamingSnakeCase}
	cuc2ss = varcaser.Caser{From: varcaser.UpperCamelCase, To: varcaser.ScreamingSnakeCase}
)

//==========================================================================
// template functions
