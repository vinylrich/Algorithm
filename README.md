# Algorithm
## TIL에서 정리했던 알고리즘을 직접 코드로 구현해보고, 문제를 풀어보는 레포지토리입니다.


1. 문제를 풀고 난 후에는 다른사람이 푼 풀이를 보자
2. 풀지 못한 문제는 다른사람의 예제를 보며 학습한다.
3. 자료구조,정렬 같은 경우는 계속 복습한다!


### Golang에서 입출력 속도를 줄이는 법

입출력 방식을 fmt.Print와 같은 방식으로 사용하면 시간초과가 날 수 있다.
그래서 bufio를 사용해야한다.

입력: bufio.NewReader(os.Stdin),scanner:= bufio.NewScanner(os.Stdin) scanner.Scan()
출력: bufio.NewWriter(os.stdout)
```