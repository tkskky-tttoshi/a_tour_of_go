package main
import(
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number:%f", float64(e))
}

func Sqrt(x float64) (float64, error){
	if x < 0{
		//Error
		//fmt.Println(ErrNegativeSqrt(x))
		return 0, ErrNegativeSqrt(x)
	}

	 // 初期値
	 z := x / 2.0
	 for i := 0; i < 10; i++ {
		 diff := (z*z - x) / (2.0 * z)
		 z = z - diff
		 if math.Abs(diff) < 1e-9 {
			 break
		 }
	 }
	 return z, nil
}

//func Sqrt(x float64) float64 {
//    z := float64(2.)
//    s := float64(0)
//    for i := 0; i < 10; i++ {
//        z = z - (z*z - x)/(2*z)
//        if math.Abs(z - s) < 1e-10 {
//            break;
//        }
//        s = z
//    }
//    return z
//}

func main(){
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

}
