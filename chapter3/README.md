## io系インタフェース実装表  
| 変数 | io.Reader | io.Writer | io.Seeker | io.Closer | io.ReaderAt |  
| -- | -- | -- | -- | -- | -- |  
| os.Stdin | ○ |  |  | ○ | ? |  
| os.File | ○ | ○ | ○ | ○ | ○ |  
| net.Conn | ○ | ○ |  | ○ |  |  

## メモリバッファが満たすインタフェース  
| 変数 | io.Reader | io.Writer | io.Seeker | io.Closer | io.ReaderAt |  
| -- | -- | -- | -- | -- | -- |  
| bytes.Buffer | ○ | ○ |  |  |  |  
| bytes.Reader | ○ |  | ○ |  | ○ |  
| strings.Reader | ○ |  | ○ |  | ○ |  
