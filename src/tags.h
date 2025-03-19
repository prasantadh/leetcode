#ifndef TAGS_H
#define TAGS_H

#include <vector>
#include <iostream>

using namespace std;

// Tags are taken from Neetcode Roadmap
enum Tags {
  ARRAYS_AND_HASHING,
  DYNAMIC_PROGRAMMING,
  TWO_POINTERS,
  STACK,
};

extern vector<Tags> tags;

// this function does not do anything. 
// It is simply to force the files that include this header
// to define tags because if tags isn't used, it doesn't need to be defined
__attribute__((constructor)) void print_tags() {
  for (auto tag: tags) {
    cout << "Tag: " << tag << endl;
  }
}

#endif // !TAGS_H
