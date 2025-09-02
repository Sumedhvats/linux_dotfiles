class DNode<T> {
    T data;
    DNode<T> next;
    DNode<T> prev;

    DNode(T val) {
        data = val;
        next = null;
        prev = null;
    }
}

class DoublyLinkedList<T> {
    private DNode<T> head;
    private DNode<T> tail;

    DoublyLinkedList() {
        head = null;
        tail = null;
    }

    void insertAtBeginning(T data) {
        DNode<T> newNode = new DNode<>(data);
        if (head == null) {
            head = newNode;
            tail = newNode;
        } else {
            newNode.next = head;
            head.prev = newNode;
            head = newNode;
        }
    }

    void insertAtEnd(T data) {
        DNode<T> newNode = new DNode<>(data);
        if (tail == null) {
            head = newNode;
            tail = newNode;
        } else {
            tail.next = newNode;
            newNode.prev = tail;
            tail = newNode;
        }
    }

    T deleteFromBeginning() {
        if (head == null) {
            System.out.println("List is empty. Nothing to delete.");
            return null;
        }
        T val = head.data;
        if (head == tail) {
            head = null;
            tail = null;
        } else {
            head = head.next;
            head.prev = null;
        }
        return val;
    }

    T deleteFromEnd() {
        if (tail == null) {
            System.out.println("List is empty. Nothing to delete.");
            return null;
        }
        T val = tail.data;
        if (head == tail) {
            head = null;
            tail = null;
        } else {
            tail = tail.prev;
            tail.next = null;
        }
        return val;
    }

    void display() {
        if (head == null) {
            System.out.println("(empty)");
            return;
        }
        DNode<T> current = head;
        while (current != null) {
            System.out.print(current.data + " <-> ");
            current = current.next;
        }
        System.out.println("NULL");
    }
}

class CNode<T> {
    T data;
    CNode<T> next;

    CNode(T val) {
        data = val;
        next = null;
    }
}

class CircularLinkedList<T> {
    private CNode<T> tail;

    CircularLinkedList() {
        tail = null;
    }

    void insertAtBeginning(T data) {
        CNode<T> newNode = new CNode<>(data);
        if (tail == null) {
            tail = newNode;
            tail.next = tail;
        } else {
            newNode.next = tail.next;
            tail.next = newNode;
        }
    }

    void insertAtEnd(T data) {
        CNode<T> newNode = new CNode<>(data);
        if (tail == null) {
            tail = newNode;
            tail.next = tail;
        } else {
            newNode.next = tail.next;
            tail.next = newNode;
            tail = newNode;
        }
    }

    T deleteFromBeginning() {
        if (tail == null) {
            System.out.println("List is empty. Nothing to delete.");
            return null;
        }
        CNode<T> temp = tail.next;
        T val = temp.data;
        if (tail.next == tail) {
            tail = null;
        } else {
            tail.next = temp.next;
        }
        return val;
    }

    T deleteFromEnd() {
        if (tail == null) {
            System.out.println("List is empty. Nothing to delete.");
            return null;
        }
        T val = tail.data;
        if (tail.next == tail) {
            tail = null;
            return val;
        }
        CNode<T> current = tail.next;
        while (current.next != tail) {
            current = current.next;
        }
        current.next = tail.next;
        tail = current;
        return val;
    }

    void display() {
        if (tail == null) {
            System.out.println("(empty)");
            return;
        }
        CNode<T> current = tail.next;
        do {
            System.out.print(current.data + " -> ");
            current = current.next;
        } while (current != tail.next);
        System.out.println("(head)");
    }
}

public class Main {
    public static void main(String[] args) {
        System.out.println("--- Doubly Linked List Operations ---");
        DoublyLinkedList<Integer> dll = new DoublyLinkedList<>();
        dll.display();
        dll.insertAtEnd(10);
        dll.insertAtEnd(20);
        dll.insertAtBeginning(5);
        dll.display();
        dll.deleteFromBeginning();
        dll.display();
        dll.deleteFromEnd();
        dll.display();

        System.out.println("\n--- Circular Linked List Operations ---");
        CircularLinkedList<String> cll = new CircularLinkedList<>();
        cll.display();
        cll.insertAtEnd("A");
        cll.insertAtEnd("B");
        cll.insertAtBeginning("X");
        cll.display();
        cll.deleteFromBeginning();
        cll.display();
        cll.deleteFromEnd();
        cll.display();
    }
}
