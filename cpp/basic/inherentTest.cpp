#include <stdio.h>

class Father
{
public:
    void foo() { }

private:
    int member1_;
    int member2_;    
};

class Son: public Father
{
public:
    void sonFoo() { }

private:
    int member1_;
    int member2_;

};

class SonWife
{
public:
    void sonWifeFoo() { }

private:
    int member1_;
    int member2_;
};

class GrandSon: public SonWife, public Son
{

public:
    void grandSonFoo() { }

private:
    int member1_;
    int member2_;

};

int main()
{
    GrandSon gsobj;
    Father* pf = &gsobj;
    Son* ps = &gsobj;
    SonWife* psw = &gsobj;
    GrandSon *pgs = &gsobj;

    printf("pf : 0x%llx\n", (unsigned long long)pf);
    printf("ps: 0x%llx v.s. psw: 0x%llx\n", (unsigned long long)ps, (unsigned long long)psw);
    printf("pgs: 0x%llx\n", (unsigned long long)pgs);

}