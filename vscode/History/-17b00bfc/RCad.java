
class Animal {
    void makeSound() {
        System.out.println("Some generic animal noise.");
    }
}

class Dog extends Animal {
    @Override
    void makeSound() {
        System.out.println("Dog says: Bow Bow");
    }
}

class Cat extends Animal {
    @Override
    void makeSound() {
        System.out.println("Cat says: Purr Meow!");
    }
}

public class Main3 {
    public static void main(String[] args) {
        Animal a1 = new Dog();
        a1.makeSound();

        Animal a2 = new Cat();
        a2.makeSound();
    }
}
