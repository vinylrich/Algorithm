# Algorithm
## 알고리즘을 공부하기 위해 만든 레포지토리입니다.


개발자는 최적화, 문제 해결을 위한 알고리즘,자료구조를 알아야 한다. 적어도 주1회 한문제씩 풀자!

1. 문제를 풀고 난 후에는 다른사람이 푼 풀이를 보자
2. 풀지 못한 문제는 다른사람의 예제를 보며 학습한다.
3. 자료구조,정렬 같은 경우는 계속 복습한다!


### Golang에서 입출력 속도를 줄이는 법

입출력 방식을 fmt.Print와 같은 방식으로 사용하면 시간초과가 날 수 있다.
그래서 bufio를 사용해야한다.

입력: bufio.NewReader(os.Stdin),scanner:= bufio.NewScanner(os.Stdin) scanner.Scan()
출력: bufio.NewWriter(os.stdout)
```