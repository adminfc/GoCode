package main
/* 
包可以区分命令空间（一个文件夹中不能有两个同名文件），也可以更好的管理项目。
go中他姨一个包，一般是创建一个文件夹，在该文件夹里面的go文件中，使用package关键字声明包名称。
通常，文件夹名称和包名称相同，并且，同一个文件下面只有一个包。
在go的安装目录下，有个src文件夹，下面就是默认的go包。D:\Program Files\Go\src\fmt\print.go

package fmt //包名称和文件夹名称一样
//导入其他依赖包
import (
	"internal/fmtsort"
	"io"
	"os"
	"reflect"
	"sync"
	"unicode/utf8"
)
*/

/* 
1.创建包的方法
创建一个名为dao的文件夹，再创建一个dao.go文件，最后在文件中声明包
package dao
import "fmt"

*/

func main() {
	
}