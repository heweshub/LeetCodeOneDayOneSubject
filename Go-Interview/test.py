

class A(object):

    def run(self):
        print("run A")

class B(A):
    
    def run(self):
        print("run B")

if __name__ == '__main__':
    a = A()
    a.run()