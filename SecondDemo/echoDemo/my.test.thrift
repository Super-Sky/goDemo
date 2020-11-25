namespace go  my.test.demo
namespace py  my.test.demo
 
struct Student{
 1: i32 sid, 
 2: string sname,
 3: bool ssex=0,
 4: i16 sage,
}
 
const map<string,string> MAPCONSTANT = {'hello':'world', 'goodnight':'moon'}
     
service ClassMember {        
    list<Student> List(1:i64 callTime),
    void Add(1: Student s),
    bool IsNameExist(1:i64 callTime, 2:string name),
}
