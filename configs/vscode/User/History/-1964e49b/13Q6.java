

abstract class Shape {
    abstract double calculateArea();
    void display() {
        System.out.println("This is a shape.");
    }
}
class Circle extends Shape {
    double radius;

    Circle(double radius) {
        this.radius = radius;
    }
    @Override
    double calculateArea() {
        return Math.PI * radius * radius;
    }
}
class Rectangle extends Shape {
    double length, width;

    // Constructor
    Rectangle(double length, double width) {
        this.length = length;
        this.width = width;
    }

    @Override
    double calculateArea() {
        return length * width;
    }
}

public class Main2 {
    public static void main(String[] args) {

        Shape circle = new Circle(8);
        circle.display();
        System.out.println("Circle Area: " + circle.calculateArea());

        System.out.println();

        Shape rectangle = new Rectangle(2, 6);
        rectangle.display();
        System.out.println("Rectangle Area: " + rectangle.calculateArea());
    }
}
