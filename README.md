# 루비로 배우는 객체지향 디자인

## 객체지향 디자인이란 무엇인가?

객체들을 협력을 어떻게 할 것인지, 어떤 역할을 가지게 할 것인지, 의존성 어떻게 가지게 할 것인지 결정하는 것.

### 객체지향

객체들이 각자 역할을 가지고 동작하도록 하여 프로그램을 만드는 것.
객체들이 협력을 할 때는 메세지로 협력한다. 그 과정에서 의존성이 생길 수 있다.

### 디자인

코드를 배치하는 것.

### 왜 객체지향 디자인을 해야할까요?

객체지향 디자인은 의존성을 관리하는 기술인데, 잘 관리하면 추후에 있을 유지보수(기능 추가, 내부 코드 변경)가 용이하다.

## 단일 책임 원칙이란 무엇인가?

하나의 책임만 있어야 한다. 클래스 안의 모든 것들이 하나의 핵심 목표와 연관되어 있을 때 강하게 응집되어 있다고 말하는데 이 것을 하나의 책임만 가지고 있다고 한다.

### 왜 단일 책임 원칙을 지켜야 하는가?

복잡해진다. 최대한 책임을 격리시켜야 한다.

### 어떻게 단일 책임 원칙을 지킬 수 있는가?

최대한 책임을 격리시켜야 한다.
어떤 객체가 여러 책임을 가지고 있다면, 분리할 수 있다.

* 인수턴수 변수 숨기기 => 여러 곳에서 참조하고 있는 데이터에서 단 한만 정의된 행동으로 바꿀 수 있다.
* 데이터 구조 숨기기 => 데이터 구조에 의존적이면, 변경에 취약해진다.
* 메서드에서 추가적인 책임을 뽑아내기 => 객체의 단일 책임 원칙을 메서드에도 확장.

## 의존성이란 무엇인가?

객체들이 협력할 때, 객체가 다른 객체에 대해 지식을 갖고 있는 상태를 의존성이라고 한다.

### 의존성이 있는지 어떻게 확인할 수 있을까요?

다른 클래스의 이름이 있나요?
다른 객체에게 보낼 메세지의 이름이 포함되어 있는가?
포함되어 있는 인자들을 알고 있는가?
인자들을 전달하는 순서를 알고 있나요?

### 의존성을 어떻게 없앨 수 있을까요?

의존성 주입 => 나한테 필요한 객체를 생성해주는 것. 생성을 바깥에서 하고 인자를 바깥에서 받는 것.
의존성을 격리 => 
의존성의 방향을 결정

=> 변화하기 쉬운 것 보다는 변화하기 어려운 것에 의존하게 하는 것.

## 인터페이스란?

객체들이 소통하는 창구, 소통할 때 사용하는 것 
메시지의 묶음.
시그니처의 모음.
* 퍼블릭 인터페이스
* 프라이빗 인터페이스

### 왜 퍼블릭 인터페이스와 프라이빗 인터페이스를 구분하는가?

다른 객체가 의존해도 되는지 안되는지를 알려주기 위해서
외부에 영향을 줄 수 있는 것들은 배제하기 위해서.
내부를 마음대로 바꾸기 위해서.

### 데메테르 원칙

의존성을 멀리 퍼트리지 말고 최대한 가까운 객체들끼리만 의존하도록 하는 것.
의존성이 많이 엮여 있을 때 어떻게 의존성을 분리할 수 있을까?
분리할 수 없을 때 어떻게 해야할까?
=> 격리시킨다.

## Duck typing이란?

