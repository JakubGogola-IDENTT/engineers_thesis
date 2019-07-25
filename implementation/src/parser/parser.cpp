#include <iostream>

#include "parser.hpp"
#include "../structures/params.hpp"

Parser::Parser() {

}

Parser::~Parser() {

}

void Parser::parse_args(char **args) {

}

/**
 * Returns parameters set
 */
std::shared_ptr<params> Parser::get_parameters() {
    return this->parameters;
}

