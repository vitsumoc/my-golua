p = person.new("Steeve")

print(person)
print(p)

print(p:name()) -- "Steeve"
p:name("Alice")
print(p:name()) -- "Alice"