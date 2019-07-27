#!/bin/ruby

require 'json'
require 'stringio'

class Circle
    attr_accessor :size
    attr_accessor :parent

    def initialize
        @size = 1
    end

    def union(other)
        r1 = self.root
        r2 = other.root
        return r1.size if r1 == r2

        r1, r2 = [r2, r1] if r2.size < r1.size
        r1.parent = r2
        r2.size += r1.size
    end

    def root
        parent ? parent.root : self
    end

    def self.get(x)
        @g ||= {}
        @g[x] ||= Circle.new
    end
end

# Complete the maxCircle function below.
def maxCircle(queries)
    m = 0
    queries.map {|a, b|
        s = Circle.get(a).union(Circle.get(b))
        m = s > m ? s : m
    }
end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

q = gets.to_i

queries = Array.new(q)

q.times do |i|
    queries[i] = gets.rstrip.split(' ').map(&:to_i)
end

ans = maxCircle queries

fptr.write ans.join "\n"
fptr.write "\n"

fptr.close()
