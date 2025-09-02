#include <iostream>
#include <vector>
using namespace std;

template <typename T>
class Stack {
private:
    int topIndex = -1;
    int capacity;
    vector<T> arr;

public:
    Stack(int s) {
        arr.reserve(s);
        capacity = s;
    }

    void push(T val) {
        if (topIndex == capacity - 1) return;
        arr.push_back(val);
        topIndex++;
    }

    bool is_empty() {
        return topIndex == -1;
    }

    void pop() {
        if (is_empty()) throw runtime_error("Stack underflow");
        arr.pop_back();
        topIndex--;
    }

    T top() {
        if (is_empty()) throw runtime_error("Stack is empty");
        return arr[topIndex];
    }

    int get_size() {
        return capacity;
    }
};

int main() {
    Stack<int> s(5);
    s.push(10);
    s.push(20);
    cout << s.top() << endl;
    s.pop();
    cout << s.top() << endl;

    Stack<string> st(3);
    st.push("A");
    st.push("B");
    cout << st.top() << endl;
}
