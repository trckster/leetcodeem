#include <iostream>
#include <vector>

using namespace std;

class NestedInteger {
private:
    bool isInt;
    int value;
    vector <NestedInteger> list;
public:
    NestedInteger(int value) {
        this->isInt = true;
        this->value = value;
    }

    NestedInteger(vector <NestedInteger> list) {
        this->isInt = false;
        this->list = list;
    }

    bool isInteger() const {
        return this->isInt;
    }

    int getInteger() const {
        return this->value;
    }

    const vector <NestedInteger> &getList() {
        return list;
    }
};


class NestedIterator {
private:
    vector <NestedInteger> nestedList;
    int index;
    NestedIterator *nestedIterator;
public:
    NestedIterator(vector <NestedInteger> &nestedList) {
        this->nestedList = nestedList;
        index = 0;
        nestedIterator = 0;
        skipEmpty();
    }

    void unsetIterator() {
        delete nestedIterator;
        nestedIterator = 0;
    }

    void skipEmpty() {
        while (index < nestedList.size()) {
            if (nestedList[index].isInteger()) {
                break;
            }

            if (!nestedIterator) {
                vector <NestedInteger> tmp = nestedList[index].getList();
                nestedIterator = new NestedIterator(tmp);
            }

            if (nestedIterator->hasNext()) {
                break;
            }

            unsetIterator();
            index++;
        }
    }

    int next() {
        int value;

        if (nestedList[index].isInteger()) {
            value = nestedList[index++].getInteger();
            skipEmpty();
        } else {
            value = nestedIterator->next();
        }

        if (nestedIterator && !nestedIterator->hasNext()) {
            unsetIterator();
            index++;
            skipEmpty();
        }

        return value;
    }

    bool hasNext() {
        if (nestedList.empty()) {
            return false;
        }

        return index < nestedList.size() || nestedIterator && nestedIterator->hasNext();
    }

    ~NestedIterator() {
        if (nestedIterator) {
            delete nestedIterator;
        }
    }
};

int main() {
    NestedInteger nf(34);
    NestedInteger nl(10);
    NestedInteger n2(5);
    NestedInteger n3(2);
    NestedInteger n4(7);
    vector <NestedInteger> n234{n2, n3, n4};
    NestedInteger nm(n234);
    vector <NestedInteger> main{nf, nm, nl};

    NestedIterator i(main);

    while (i.hasNext()) {
        cout << i.next() << endl;
    }
}