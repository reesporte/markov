'''basic markov chain to generate tai lopez quotes'''
from string import punctuation
from random import choice, randint
from time import sleep

NGRAM_LEVEL = 40

def tokenize(words: str):
    '''split into char level tokens'''
    return [c for c in words if not c == "\n"]

def ngrams(words: str):
    '''generate ngrams from a string'''
    words = tokenize(words)
    ngrams = {}
    for i in range(len(words) - NGRAM_LEVEL):
        if not ngrams.get(words[i], False):
            ngrams[words[i]] = []
        for j in range(1, NGRAM_LEVEL):
            ngrams[words[i]].append(''.join(words[i + 1: i + 1 + j]))

    return words, ngrams


def create_quote(tokens: list, grams: dict):
    '''generate quote from list of tokens and associated ngrams'''
    last_word = choice(tokens)
    quote = last_word.capitalize()
    for i in range(randint(30, 45)):
        last_word = choice(grams[last_word[-1]])
        quote += last_word
    return quote


def main():
    '''read words and generate quote'''
    with open('words.txt', 'r') as f:
        words = f.read()
    tokens, grams = ngrams(words)
    print(create_quote(tokens, grams))

if __name__ == '__main__':
    main()
