

class Animal {
    void sound() {
        System.out.println("Animal makes a sound.");
    }
}
class Dog extends Animal {
    @Override
    void sound() {
        System.out.println("Dog barks: Woof!");
    }
}
class Cat extends Animal {
    @Override
    void sound() {
        System.out.println("Cat meows: Meow!");
    }
}
public class Main3 {
    public static void main(String[] args) {
        Animal dog = new Dog();
        dog.sound();
        Animal cat = new Cat();
        cat.sound();
    }
}
