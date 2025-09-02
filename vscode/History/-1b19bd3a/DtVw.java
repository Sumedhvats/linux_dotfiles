
interface Driveable {
    void drive();
}

class Vehicle {
    String company;
    int maxSpeed;

    Vehicle(String company, int maxSpeed) {
        this.company = company;
        this.maxSpeed = maxSpeed;
    }

    void move() {
        System.out.println(company + " is moving at " + maxSpeed + " km/h.");
    }
}

class Car extends Vehicle implements Driveable {
    int seatCount;

    Car(String company, int maxSpeed, int seatCount) {
        super(company, maxSpeed);
        this.seatCount = seatCount;
    }

    @Override
    public void drive() {
        System.out.println("Driving smoothly and comfortably.");
    }

    void showDetails() {
        System.out.println("Car Company: " + company);
        System.out.println("Top Speed: " + maxSpeed + " km/h");
        System.out.println("Number of Seats: " + seatCount);
    }
}

public class Main1 {
    public static void main(String[] args) {
        Car car1 = new Car("Range Rover", 200, 5);
        Car car2 = new Car("Defender", 150, 7);

        car1.move();
        car1.drive();
        car1.showDetails();

        System.out.println();

        car2.move();
        car2.drive();
        car2.showDetails();
    }
}
