# ERROR的错误处理注意事项

可以使用github.com/pkg/errors包来代替标准库的errors包, 该包可以通过`errors.Wrap(f)`和`errors.WithMessage`来添加上下文信息

**注:**  `errors.Wrap(f)`保存了error的堆栈信息, `errors.WithMessage`不保存堆栈信息. 

1. 首次产生错误的时候对error进行warp 
   - **注:**  `首次`是指在业务代码与标准库, 第三方库等交互以及自己的代码**生成**错误时. 自己的代码中生成错误信息使用`errors.New`或`errors.Errorf`

2. Error只处理1一次, 打日志(最顶层)或者向上返回, 不要操作两次, 打印日志可以通过使用`%+v`谓词把堆栈信息打印出来.
   - 如果不处理error, 那么要使用Wrap(f)或WithMessage添加一些上下文信息向上返回. **注:** 不需要把整个response打印出来, 因为它的内容太多
   - 如果处理了error(打日志 降级 or 其他逻辑), 那么不要再将error向上返回

3. 如果在自己的代码中调用其他函数, 要直接返回, 不要再次进行warp, 但是可以调用WithMessage添加上下文信息.
4. 可以使用`errors.Cause`获取根因来与`sentinel error`进行判定
5. kit(基础库)或标准库不应该对error进行Wrap, 只能返回根因, 业务层代码可以进行Wrap. 
6. error的Unwrap方法可以将根因返回
7. go1.13新特性 `errors.Is(结合errors.Cause与sentinel error进行判断)` 和 `errors.As(起一个断言的作用)` 的使用, 不要再直接与sentinel error进行等值判断
   - `errors.Is` 会与sentinel error进行等值判断, 如果判断没过, 则尝试通过Unwrap获取error的根因
   - `errors.As` 可以将err对象转换为自定义的error类型










