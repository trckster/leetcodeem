#include<iostream>
#include <queue>

using namespace std;

class MyStack {
private:
    queue<int> main, secondary;
public:
    MyStack() {

    }

    void push(int x) {
        while (!main.empty()) {
            secondary.push(main.front());
            main.pop();
        }

        main.push(x);

        while (!secondary.empty()) {
            main.push(secondary.front());
            secondary.pop();
        }
    }

    int pop() {
        int result = main.front();

        main.pop();

        return result;
    }

    int top() {
        return main.front();
    }

    bool empty() {
        return main.empty();
    }
};

int main() {
    int x = 5;

    MyStack *obj = new MyStack();
    obj->push(x);
    int param_3 = obj->top();
    cout << (obj->empty() ? 'E' : 'N') << endl;
    int param_2 = obj->pop();
    cout << (obj->empty() ? 'E' : 'N') << endl;

    delete obj;
}