#include <iostream>
#include <fstream>

using namespace std;

int getNum(string line) {
    int num = 0;

    for(char c : line) {
        if(isdigit(c)) {
            num += c - '0';
            break;
        }
    }

    num *= 10;

    for(int i = line.size() - 1; i >= 0; i--) {
        if(isdigit(line[i])) {
            num += line[i] - '0';
            break;
        }
    }

    return num;
}

int main() {
    ifstream file("day01.txt");
    string line;
    int ans = 0;
    while(getline(file, line)) {
        ans += getNum(line);
    }
    cout << ans << "\n";
}
