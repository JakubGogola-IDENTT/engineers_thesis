#pragma once
#include <memory>

#include "../structures/params.hpp"

class Parser {
private:
    std::shared_ptr<params> parameters;

public:
    Parser() {}
    ~Parser() {}

    void parse_args(char **args);
    std::shared_ptr<params> get_parameters();
};
