#include<iostream>

using namespace std;

class Node {
public:
    int value;
    Node *next;

    Node(int value) {
        this->value = value;
        this->next = 0;
    }

    bool has(int value) {
        if (this->value == value) {
            return true;
        }

        return this->next && this->next->has(value);
    }

    void add(int value) {
        if (this->next) {
            this->next->add(value);
        } else {
            this->next = new Node(value);
        }
    }

    Node *remove(int toBeRemoved) {
        if (this->next) {
            this->next = this->next->remove(toBeRemoved);
        }

        if (this->value != toBeRemoved) {
            return this;
        }

        Node *toReturn = this->next;

        this->next = 0;
        delete this;

        return toReturn;
    }

    ~Node() {
        if (this->next) {
            delete this->next;
        }
    }
};

const int containerLength = 10001;

class MyHashSet {
public:
    Node *roots[containerLength];

    MyHashSet() {
        for (int i = 0; i < containerLength; ++i) {
            roots[i] = 0;
        }
    }

    int getIndex(int key) {
        return key % containerLength;
    }

    void add(int key) {
        int index = getIndex(key);

        if (!roots[index]) {
            roots[index] = new Node(key);
        } else {
            if (!roots[index]->has(key)) {
                roots[index]->add(key);
            }
        }
    }

    void remove(int key) {
        int index = getIndex(key);

        if (!roots[index]) {
            return;
        }

        roots[index] = roots[index]->remove(key);
    }

    bool contains(int key) {
        int index = getIndex(key);

        if (!roots[index]) {
            return false;
        }

        return roots[index]->has(key);
    }

    ~MyHashSet() {
        for (int i = 0; i < containerLength; ++i) {
            if (roots[i]) {
                delete roots[i];
            }
        }
    }
};

int main() {
    MyHashSet *obj = new MyHashSet();
    obj->add(5);
    obj->add(10006);
    obj->add(5);
    string res = obj->contains(5) ? "true" : "false";
    cout << res << endl;
    res = obj->contains(10006) ? "true" : "false";
    cout << res << endl;
    obj->remove(5);
    res = obj->contains(5) ? "true" : "false";
    cout << res << endl;
    res = obj->contains(10006) ? "true" : "false";
    cout << res << endl;

    delete obj;
}