
class BankAccount {
    private String accNo;
    private double currentBalance;

    BankAccount(String accNo, double initialBalance) {
        this.accNo = accNo;
        if (initialBalance >= 0) {
            this.currentBalance = initialBalance;
        } else {
            this.currentBalance = 0;
        }
    }

    public double getBalance() {
        return currentBalance;
    }

    
    public void setBalance(double amount) {
        if (amount >= 0) {
            this.currentBalance = amount;
        } else {
            System.out.println("Balance cannot be negative.");
        }
    }


    public void deposit(double amount) {
        if (amount > 0) {
            currentBalance += amount;
            System.out.println("✅ Added " + amount + " to account.");
        } else {
            System.out.println("Deposit must be greater than zero.");
        }
    }


    public void withdraw(double amount) {
        if (amount > 0 && amount <= currentBalance) {
            currentBalance -= amount;
            System.out.println("Withdrawn: " + amount);
        } else {
            System.out.println("Withdrawal denied (invalid or insufficient funds).");
        }
    }
}
public class Main4 {
    public static void main(String[] args) {
        BankAccount myAcc = new BankAccount("ID2025", 1000);

        myAcc.deposit(400);
        myAcc.withdraw(250);

        System.out.println("Remaining Balance: " + myAcc.getBalance());
    }
}
