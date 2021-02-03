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

수리점

수리
 - 분해
 - 수리
 if (paintable()) {
   칠하다
 }
 - 조립

책상
 - 분해
 - 수리
 - 칠하다(Nothing)
 - 조립
 @override
 paintable() {
   return false;
 }

카메라
 - 분해
 - 수리
 @overide
 - 칠하다 (기름칠)
  
 - 조립

자전거
 - 분해
 - 수리
 - 기름칠
 - 조립
노트북
 - 분해
 - 수리
 - 써멀구리스바르기
 - 조립
키보드
 - 분해
 - 수리
 - 윤활유바르기
 - 조립

```typescript
BadRequest
InternalServerError
ProductNotFoundError

handleError(err Error) {
  show(err.message);

  if (err.data) {
    // doSomething
  }
}

catch (err Error) {
  handleError(err);
  // if (err instanceof BadRequest) {
  //   show(err.message);
  // }

  // if (err instanceof ProductNotFoundError) {
  //   show(err.message);
  //   products.map(doSomething);
  // }

  // if (err instanceof InternalServerError) {

  // }
}
```

## 상속이란?

자동화된 메시지 전달 시스템이다. 상위 클래스는 일반화된 행동을 가지고 있고, 하위 클래스는 구체적 행동을 가지고 있다.

### 리스코프 취환 원칙

하위클래스가 상위클래스를 대체할 수 있다.

```javascript
class Studuent {
  getGrade(): number;
}

const grade: Grade = student.getGrade();

class ElementaryStudent extends Studuent {
  getGrade() {
    return 1 ~ 6;
  }
}

class MiddleStduent extends Studuent {
  getGrade() {
    return 1 ~ 3;
  }
}

class Duck {
  fly() {
    return 날개;
  }
}

RubberDuck.fly()

class RubberDuck {
  fly() {
    return null;
  }
}
```

## 오리 타입 이란?

요구하는 행동을 구현했다면 그 타입으로 취급하는 것.

## 조합이란?

독립적인 객체를 보다 크고 복합적인 것으로 통합하는 것

```javascript
class Coffee() {
  make() {

  }

  boil() {

  }
}

class FruitAde() {
  make() {
  }

  shake() {

  }
}

class 스무디() {
  make() {

  }

  shake() {

  }
}

class 흑당밀크티() {
  make() {

  }
}

class 달고나커피() {
  make() {
    
  }

  shake() {

  }
}
```

```javascript
class Beberage() {
  make() {}
  shake() {

  }
  boil() {

  }
}

class Coffee() extends Beberage {
  make() {}
  shake() {
    // Do nothing
  }
  boil() {}
}

class 스무디() extends Beberage {
  make() {}
  shake() {}
  boil() {
    // Do nothing
  }
}
```

```javascript
interface Boilable {
  boil();
}

interface Shakable {
  shake();
}

class Coffee() {
}

class 스무디() extends Beberage implements Shakable {
  shake() {}
}

Coffee

Non Coffee

Beberage.shake();
```

is-a => 상속써라
아메리카노는 커피다.
디카페인 아메리카노는 커피다.

behaves-like-a 오리 타입
아메리카노 implements Boilable
레몬에이드

아메리카노는.boil();
레몬에이드.boil(); (X)

아메라카노는 끓을 수 있다.
디카페인 아메리카노도 끓을 수 있다.
차도 끓을 수 있다.

has-a => 조합써라
커피는 시럽을 가지고 있다.
커피는 첨가물 가지고 있다.

## 조합이란?

## 상속과 조합

## 테스트

### private 메서드를 테스트를 안해야하는 이유는?

자주 변경되어서.
애플리케이션 관점에서는 프라이빗 메서드는 없는 것이다.

### 수정하기 쉬운 코드를 작성하는데 필요한 세가지는?

1. 객체지향 디자인을 이해해햔다.
2. 리팩터링을 해야한다.
3. 테스트를 작성할 수 있어야 한다.

### 테스트를 짜야하는 이유는?

변경을 했을 때, 우리의 코드에 어떤 영향을 주는지 테스트가 확인해준다.
문서를 제공해준다. 테스트를 통해서 어떻게 동작하는 기술할 수 있다.
고장났을 때 어디가 고장났는지 빠르게 확인할 수 있다. 그래서 비용을 줄여준다.
코드를 빠르게 사용해볼 수 있다.
디자인의 결점을 발견할 수 있다.
추상화를 돕기
결정을 미룰 수 있다.
