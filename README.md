
# go实现 chatbot 问答

## 测试结果

```go
func TestAsk(t *testing.T) {
	question := "使用 java 写一个冒泡排序"
	repository := impl.NewOpenAIRepository()
	gtp := repository.DoChatGTP(question)
	fmt.Println(gtp)
}
```

```bash
=== RUN   TestAsk


```java
public class BubbleSort {
    public static void main(String[] args) {
        int[] arr = {3, 9, -1, 10, -2};

        // 外层循环控制排序趟数
        for (int i = 0; i < arr.length - 1; i++) {
            // 内层循环控制每一趟排序多少次
            for (int j = 0; j < arr.length - 1 - i; j++) {
                if (arr[j] > arr[j + 1]) {
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                }
            }
        }

        System.out.println("排序后的数组：");
        for (int i = 0; i < arr.length; i++) {
            System.out.print(arr[i] + " ");
        }
    }
}
```
--- PASS: TestAsk (21.77s)
PASS
```