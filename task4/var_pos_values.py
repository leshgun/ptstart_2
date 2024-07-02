import re



def find_int_values(text: str) -> list:
    # Get content of method of the class
    cl = re.search(r'class[^{]*{([\s\S]*)}', text)[1]
    me = re.search(r'method[^{]*{([\s\S]*)}', cl)[1]

    # Get the name of variable
    var = re.search(r"int ([^;]+);", me)[1]
    
    # Normilize source code
    text = ''.join([l.strip() for l in me.split('\n') if l])

    # Count oppened/closed brackets
    bra_open = 0
    bra_closed = 0
    bra_is_open = False
    values = {}
    i = 1
    while i < len(text):
        # find "x = *value*" statements at start of string
        m = re.match(rf'{var}\s*=\s*(\d*)', text[i:])
        if m:
            # m[1] = value of statement
            if bra_is_open:
                values[bra_open] = m[1]
            else:
                # if statement without "if" then override values of current scope 
                # of the variable and above
                for j in range(bra_open - bra_closed, bra_open + 1):
                    values[j] = m[1]
            # escape the matched string
            i += m.end()
            continue
        if text[i] == '}':
            bra_closed += 1
            i += 1
            bra_is_open = False
            continue
        m = re.match(r'if[^{]*{', text[i:])
        if m:
            bra_open += 1
            # escape the matched string
            i += m.end()
            bra_is_open = True
            continue
        i += 1

    print(sorted(list(set(values.values()))))

def main():
    with open("Test.java") as file:
        find_int_values(file.read())
        file.close()
        
if __name__=='__main__':
    main()
