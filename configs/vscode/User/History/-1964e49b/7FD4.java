
abstract class Shape {
    abstract double findArea();

    void show() {
        System.out.println("Generic shape information.");
    }
}
class Circle extends Shape {
    double r;

    Circle(double r) {
        this.r = r;
    }

    @Override
    double findArea() {
        return Math.PI * r * r;
    }
}

class Rectangle extends Shape {
    double l, w;

    Rectangle(double l, double w) {
        this.l = l;
        this.w = w;
    }

    @Override
    double findArea() {
        return l * w;
    }
}

public class Main2 {
    public static void main(String[] args) {

        Shape c = new Circle(8);
        c.show();
        System.out.println("Area of Circle: " + c.findArea());

        System.out.println();

        Shape r = new Rectangle(2, 6);
        r.show();
        System.out.println("Area of Rectangle: " + r.findArea());
    }
}
