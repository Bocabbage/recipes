#include <iostream>
#include <stdio.h>
using namespace std;

class Base {
public:
	virtual void a() { cout << "Base a()" << endl; }
	virtual void b() { cout << "Base b()" << endl; }
	virtual void c() { cout << "Base c()" << endl; }
};

class Derive : public Base {
public:
	virtual void b() { cout << "Derive b()" << endl; }
};

int main()
{
	typedef void (*Func)();
	cout << "-----------Base------------" << endl;
	Base* q = new Base;
	long* tmp1 = (long*)q;
	long* vptr1 = (long*)(*tmp1);
	for (int i = 0; i < 3; i++) {
		printf("vptr[%d] : %p\n", i, vptr1[i]);
	}
	Func a = (Func)vptr1[0];
	Func b = (Func)vptr1[1];
	Func c = (Func)vptr1[2];
	a();
	b();
	c();

	Derive* p = new Derive;
	long* tmp = (long*)p;
	long* vptr = (long*)(*tmp);
	cout << "---------Derive------------" << endl;
	for (int i = 0; i < 3; i++) {
		printf("vptr[%d] : %p\n", i, vptr[i]);
	}
	Func d = (Func)vptr[0];
	Func e = (Func)vptr[1]; // 与基类应当不同：为多态派生的void b()
	Func f = (Func)vptr[2];
	d();
	e();
	f();

	
	return 0;
}