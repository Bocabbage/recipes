#pragma once

class Singleton
{
public:
    ~Singleton() = default;
    static Singleton& getInstance()
    {
        // initialized the first time control passes through their declaration
        static Singleton uniqueInstance_;
        return uniqueInstance_;
    }

private:
    Singleton() = default;
    Singleton(const Singleton&) = delete;
    Singleton& operator()(const Singleton&) = delete;

};
