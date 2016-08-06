package main

// #include "Main.h"
//
// #include "bridge.h"
//
// // JNIEnv, jclass
// #include <jni.h>
//
// // free()
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

//export Java_Main_test
func Java_Main_test(env *C.JNIEnv, clazz C.jclass) {
	fmt.Println(GetVersion(env))
	fmt.Println("Hello, world from a Go function!")
	test2(env, clazz)
}

func test2(env *C.JNIEnv, clazz C.jclass) {
	WithCString("test2", func(c_name *C.char) {
		WithCString("()V", func(c_sig *C.char) {
			if mid := C._GoJniGetStaticMethodID(env, clazz, c_name, c_sig); mid != nil {
				C._GoJniCallStaticVoidMethod(env, clazz, mid)
			}
		})
	})
}

func main() {
}

// jni.h:
//     /*
//      * JNI Types
//      */

// jni.h:
//     typedef jint jsize;
/* No need to do anything */

// jni.h:
//     struct _jobject;
//     typedef struct _jobject *jobject;
type JObject struct {
	peer C.jobject
}

func NewJObject(peer C.jobject) JObject {
	return JObject{peer}
}

func (obj JObject) Peer() C.jobject {
	return obj.peer
}

// jni.h:
//     typedef jobject jclass;
type JClass JObject

func NewJClass(peer C.jclass) JClass {
	return JClass{C.jobject(peer)}
}

func (clazz JClass) Peer() C.jclass {
	return C.jclass(clazz.peer)
}

// jni.h:
//     typedef jobject jthrowable;
type JThrowable JObject

func NewJThrowable(peer C.jthrowable) JThrowable {
	return JThrowable{C.jobject(peer)}
}

func (obj JThrowable) Peer() C.jthrowable {
	return C.jthrowable(obj.peer)
}

// jni.h:
//     typedef jobject jstring;
// type JString JObject

// jni.h:
//     typedef jobject jarray;
// type JArray JObject

// jni.h:
//     typedef jarray jbooleanArray;
// type JBooleanArray JArray

// jni.h:
//     typedef jarray jbyteArray;
// type JByteArray JArray

// jni.h:
//     typedef jarray jcharArray;
// type JCharArray JArray

// jni.h:
//     typedef jarray jshortArray;
// type JShortArray JArray

// jni.h:
//     typedef jarray jintArray;
// type JIntArray JArray

// jni.h:
//     typedef jarray jlongArray;
// type JLongArray JArray

// jni.h:
//     typedef jarray jfloatArray;
// type JFloatArray JArray

// jni.h:
//     typedef jarray jdoubleArray;
// type JDoubleArray JArray

// jni.h:
//     typedef jarray jobjectArray;
// type JObjectArray JArray

// jni.h
//     typedef jobject jweak;
type JWeak JObject

func NewJWeak(peer C.jweak) JWeak {
	return JWeak{C.jobject(peer)}
}

func (ref JWeak) Peer() C.jweak {
	return C.jweak(ref.peer)
}

// jni.h
//     typedef union jvalue {
//         jboolean z;
//         jbyte    b;
//         jchar    c;
//         jshort   s;
//         jint     i;
//         jlong    j;
//         jfloat   f;
//         jdouble  d;
//         jobject  l;
//     } jvalue;
/* Not Implemented */

// jni.h:
//     struct _jfieldID;
//     typedef struct _jfieldID *jfieldID;
type JFieldID struct {
	peer C.jfieldID
}

func NewJFieldID(peer C.jfieldID) JFieldID {
	return JFieldID{peer}
}

func (fieldID JFieldID) Peer() C.jfieldID {
	return fieldID.peer
}

// jni.h:
//     struct _jmethodID;
//     typedef struct _jmethodID *jmethodID;
type JMethodID struct {
	peer C.jmethodID
}

func NewJMethodID(peer C.jmethodID) JMethodID {
	return JMethodID{peer}
}

func (methodID JMethodID) Peer() C.jmethodID {
	return methodID.peer
}

// jni.h:
//     /* Return values from jobjectRefType */
//     typedef enum _jobjectType {
//         JNIInvalidRefType    = 0,
//         JNILocalRefType      = 1,
//         JNIGlobalRefType     = 2,
//         JNIWeakGlobalRefType = 3
//     } jobjectRefType;
const (
	JNIInvalidRefType    = C.JNIInvalidRefType
	JNILocalRefType      = C.JNILocalRefType
	JNIGlobalRefType     = C.JNIGlobalRefType
	JNIWeakGlobalRefType = C.JNIWeakGlobalRefType
)

// jni.h:
//     /*
//      * jboolean constants
//      */
//
//     #define JNI_FALSE 0
//     #define JNI_TRUE 1
const (
	JNI_FALSE = 0
	JNI_TRUE  = 1
)

// jni.h:
//     /*
//      * possible return values for JNI functions.
//      */
//
//     #define JNI_OK           0                 /* success */
//     #define JNI_ERR          (-1)              /* unknown error */
//     #define JNI_EDETACHED    (-2)              /* thread detached from the VM */
//     #define JNI_EVERSION     (-3)              /* JNI version error */
//     #define JNI_ENOMEM       (-4)              /* not enough memory */
//     #define JNI_EEXIST       (-5)              /* VM already created */
//     #define JNI_EINVAL       (-6)              /* invalid arguments */
const (
	JNI_OK        = 0
	JNI_ERR       = -1
	JNI_EDETACHED = -2
	JNI_EVERSION  = -3
	JNI_ENOMEM    = -4
	JNI_EEXIST    = -5
	JNI_EINVAL    = -6
)

// jni.h:
//     /*
//      * used in ReleaseScalarArrayElements
//      */
//
//     #define JNI_COMMIT 1
//     #define JNI_ABORT 2
const (
	JNI_COMMIT = 1
	JNI_ABORT  = 2
)

// jni.h:
//     /*
//      * used in RegisterNatives to describe native method name, signature,
//      * and function pointer.
//      */
//
//     typedef struct {
//         char *name;
//         char *signature;
//         void *fnPtr;
//     } JNINativeMethod;
/* Not Implemented */

// jni.h:
//     /*
//      * JNI Native Method Interface.
//      */
//
//     struct JNINativeInterface_;
//     typedef const struct JNINativeInterface_ *JNIEnv;
type JNIEnv struct {
	peer *C.JNIEnv
}

func NewJNIEnv(peer *C.JNIEnv) JNIEnv {
	return JNIEnv{peer}
}

func (env JNIEnv) Peer() *C.JNIEnv {
	return env.peer
}

// jni.h:
//     /*
//      * JNI Invocation Interface.
//      */
//
//     struct JNIInvokeInterface_;
//     typedef const struct JNIInvokeInterface_ *JavaVM;
type JavaVM struct {
	peer *C.JavaVM
}

func NewJavaVM(peer *C.JavaVM) JavaVM {
	return JavaVM{peer}
}

func (vm JavaVM) Peer() *C.JavaVM {
	return vm.peer
}

// jni.h:
//     jint (JNICALL *GetVersion)(JNIEnv *env);
func GetVersion(env *C.JNIEnv) C.jint {
	return C._GoJniGetVersion(env)
}

func (env JNIEnv) GetVersion() int32 {
	return int32(GetVersion(env.Peer()))
}

// jni.h:
//     jclass (JNICALL *DefineClass)(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len);
func DefineClass(env *C.JNIEnv, name *C.char, loader C.jobject, buf *C.jbyte, _len C.jsize) C.jclass {
	return C._GoJniDefineClass(env, name, loader, buf, _len)
}

func (env JNIEnv) DefineClass(name string, loader JObject, buf []byte, _len int32) JClass {
	// TODO: Implement this!
	panic("Not Yet Implemented")
}

// jni.h:
//     jclass (JNICALL *FindClass)(JNIEnv *env, const char *name);
func FindClass(env *C.JNIEnv, name *C.char) C.jclass {
	return C._GoJniFindClass(env, name)
}

func (env JNIEnv) FindClass(name string) JClass {
	var result C.jclass
	WithCString(name, func(c_name *C.char) {
		result = FindClass(env.Peer(), c_name)
	})
	return NewJClass(result)
}

// jni.h:
//     jmethodID (JNICALL *FromReflectedMethod)(JNIEnv *env, jobject method);
func FromReflectedMethod(env *C.JNIEnv, method C.jobject) C.jmethodID {
	return C._GoJniFromReflectedMethod(env, method)
}

func (env JNIEnv) FromReflectedMethod(method JObject) JMethodID {
	return NewJMethodID(FromReflectedMethod(env.Peer(), method.Peer()))
}

// jni.h:
//     jfieldID (JNICALL *FromReflectedField)(JNIEnv *env, jobject field);
func FromReflectedField(env *C.JNIEnv, field C.jobject) C.jfieldID {
	return C._GoJniFromReflectedField(env, field)
}

func (env JNIEnv) FromReflectedField(field JObject) JFieldID {
	return NewJFieldID(FromReflectedField(env.Peer(), field.Peer()))
}

// jni.h:
//     jobject (JNICALL *ToReflectedMethod)(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);
func ToReflectedMethod(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, isStatic C.jboolean) C.jobject {
	return C._GoJniToReflectedMethod(env, cls, methodID, isStatic)
}

func (env JNIEnv) ToReflectedMethod(cls JClass, methodID JMethodID, isStatic bool) JObject {
	return NewJObject(ToReflectedMethod(env.Peer(), cls.Peer(), methodID.Peer(), JavaBoolean(isStatic)))
}

// jni.h:
//     jclass (JNICALL *GetSuperclass)(JNIEnv *env, jclass sub);
func GetSuperclass(env *C.JNIEnv, sub C.jclass) C.jclass {
	return C._GoJniGetSuperclass(env, sub)
}

func (env JNIEnv) GetSuperclass(sub JClass) JClass {
	return NewJClass(GetSuperclass(env.Peer(), sub.Peer()))
}

// jni.h:
//     jboolean (JNICALL *IsAssignableFrom)(JNIEnv *env, jclass sub, jclass sup);
func IsAssignableFrom(env *C.JNIEnv, sub, sup C.jclass) C.jboolean {
	return C._GoJniIsAssignableFrom(env, sub, sup)
}

func (env JNIEnv) IsAssignableFrom(sub, sup JClass) bool {
	return GoBool(IsAssignableFrom(env.Peer(), sub.Peer(), sup.Peer()))
}

// jni.h:
//     jobject (JNICALL *ToReflectedField)(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);
func ToReflectedField(env *C.JNIEnv, cls C.jclass, fieldID C.jfieldID, isStatic C.jboolean) C.jobject {
	return C._GoJniToReflectedField(env, cls, fieldID, isStatic)
}

func (env JNIEnv) ToReflectedField(cls JClass, fieldID JFieldID, isStatic bool) JObject {
	return NewJObject(ToReflectedField(env.Peer(), cls.Peer(), fieldID.Peer(), JavaBoolean(isStatic)))
}

// jni.h:
//     jint (JNICALL *Throw)(JNIEnv *env, jthrowable obj);
func Throw(env *C.JNIEnv, obj C.jthrowable) C.jint {
	return C._GoJniThrow(env, obj)
}

func (env JNIEnv) Throw(obj JThrowable) int32 {
	return int32(Throw(env.Peer(), obj.Peer()))
}

// jni.h:
//     jint (JNICALL *ThrowNew)(JNIEnv *env, jclass clazz, const char *msg);
func ThrowNew(env *C.JNIEnv, clazz C.jclass, msg *C.char) C.jint {
	return C._GoJniThrowNew(env, clazz, msg)
}

func (env JNIEnv) ThrowNew(clazz JClass, msg string) int32 {
	var result C.jint
	WithCString(msg, func(c_msg *C.char) {
		result = ThrowNew(env.Peer(), clazz.Peer(), c_msg)
	})
	return int32(result)
}

// jni.h:
//     jthrowable (JNICALL *ExceptionOccurred)(JNIEnv *env);
func ExceptionOccurred(env *C.JNIEnv) C.jthrowable {
	return C._GoJniExceptionOccurred(env)
}

func (env JNIEnv) ExceptionOccurred() JThrowable {
	return NewJThrowable(ExceptionOccurred(env.Peer()))
}

// jni.h:
//     void (JNICALL *ExceptionDescribe)(JNIEnv *env);
func ExceptionDescribe(env *C.JNIEnv) {
	C._GoJniExceptionDescribe(env)
}

func (env JNIEnv) ExceptionDescribe() {
	ExceptionDescribe(env.Peer())
}

// jni.h:
//     void (JNICALL *ExceptionClear)(JNIEnv *env);
func ExceptionClear(env *C.JNIEnv) {
	C._GoJniExceptionClear(env)
}

func (env JNIEnv) ExceptionClear() {
	ExceptionClear(env.Peer())
}

// jni.h:
//     void (JNICALL *FatalError)(JNIEnv *env, const char *msg);
func FatalError(env *C.JNIEnv, msg *C.char) {
	C._GoJniFatalError(env, msg)
}

func (env JNIEnv) FatalError(msg string) {
	WithCString(msg, func(c_msg *C.char) {
		FatalError(env.Peer(), c_msg)
	})
}

// jni.h:
//     jint (JNICALL *PushLocalFrame)(JNIEnv *env, jint capacity);
func PushLocalFrame(env *C.JNIEnv, capacity C.jint) C.jint {
	return C._GoJniPushLocalFrame(env, capacity)
}

func (env JNIEnv) PushLocalFrame(capacity int32) int32 {
	return int32(PushLocalFrame(env.Peer(), C.jint(capacity)))
}

// jni.h:
//     jobject (JNICALL *PopLocalFrame)(JNIEnv *env, jobject result);
func PopLocalFrame(env *C.JNIEnv, result C.jobject) C.jobject {
	return C._GoJniPopLocalFrame(env, result)
}

func (env JNIEnv) PopLocalFrame(result JObject) JObject {
	return NewJObject(PopLocalFrame(env.Peer(), result.Peer()))
}

// jni.h:
//     jobject (JNICALL *NewGlobalRef)(JNIEnv *env, jobject lobj);
func NewGlobalRef(env *C.JNIEnv, lobj C.jobject) C.jobject {
	return C._GoJniNewGlobalRef(env, lobj)
}

func (env JNIEnv) NewGlobalRef(lobj JObject) JObject {
	return NewJObject(NewGlobalRef(env.Peer(), lobj.Peer()))
}

// jni.h:
//     void (JNICALL *DeleteGlobalRef)(JNIEnv *env, jobject gref);
func DeleteGlobalRef(env *C.JNIEnv, gref C.jobject) {
	C._GoJniDeleteGlobalRef(env, gref)
}

func (env JNIEnv) DeleteGlobalRef(gref JObject) {
	DeleteGlobalRef(env.Peer(), gref.Peer())
}

// jni.h:
//     void (JNICALL *DeleteLocalRef)(JNIEnv *env, jobject obj);
func DeleteLocalRef(env *C.JNIEnv, obj C.jobject) {
	C._GoJniDeleteLocalRef(env, obj)
}

func (env JNIEnv) DeleteLocalRef(obj JObject) {
	DeleteLocalRef(env.Peer(), obj.Peer())
}

// jni.h:
//     jboolean (JNICALL *IsSameObject)(JNIEnv *env, jobject obj1, jobject obj2);
func IsSameObject(env *C.JNIEnv, obj1, obj2 C.jobject) C.jboolean {
	return C._GoJniIsSameObject(env, obj1, obj2)
}

func (env JNIEnv) IsSameObject(obj1, obj2 JObject) bool {
	return GoBool(IsSameObject(env.Peer(), obj1.Peer(), obj2.Peer()))
}

// jni.h:
//     jobject (JNICALL *NewLocalRef)(JNIEnv *env, jobject ref);
func NewLocalRef(env *C.JNIEnv, ref C.jobject) C.jobject {
	return C._GoJniNewLocalRef(env, ref)
}

func (env JNIEnv) NewLocalRef(ref JObject) JObject {
	return NewJObject(NewLocalRef(env.Peer(), ref.Peer()))
}

// jni.h:
//     jint (JNICALL *EnsureLocalCapacity)(JNIEnv *env, jint capacity);
func EnsureLocalCapacity(env *C.JNIEnv, capacity C.jint) C.jint {
	return C._GoJniEnsureLocalCapacity(env, capacity)
}

func (env JNIEnv) EnsureLocalCapacity(capacity int32) int32 {
	return int32(EnsureLocalCapacity(env.Peer(), C.jint(capacity)))
}

// jni.h:
//     jobject (JNICALL *AllocObject)(JNIEnv *env, jclass clazz);
func AllocObject(env *C.JNIEnv, clazz C.jclass) C.jobject {
	return C._GoJniAllocObject(env, clazz)
}

func (env JNIEnv) AllocObject(clazz JClass) JObject {
	return NewJObject(AllocObject(env.Peer(), clazz.Peer()))
}

// jni.h:
//     jobject (JNICALL *NewObject)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *NewObjectV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *NewObjectA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func NewObjectA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {
	return C._GoJniNewObjectA(env, clazz, methodID, args)
}

func (env JNIEnv) NewObject(clazz JClass, methodID JMethodID, args []interface{}) JObject {
	// TODO: Implement this!
	panic("Not Yet Implemented")
}

// jni.h:
//     jclass (JNICALL *GetObjectClass)(JNIEnv *env, jobject obj);
func GetObjectClass(env *C.JNIEnv, obj C.jobject) C.jclass {
	return C._GoJniGetObjectClass(env, obj)
}

func (env JNIEnv) GetObjectClass(obj JObject) JClass {
	return NewJClass(GetObjectClass(env.Peer(), obj.Peer()))
}

// TODO: Here!!!

// jni.h:
//     jboolean (JNICALL *IsInstanceOf)(JNIEnv *env, jobject obj, jclass clazz);
func IsInstanceOf(env *C.JNIEnv, obj C.jobject, clazz C.jclass) C.jboolean {
	return C._GoJniIsInstanceOf(env, obj, clazz)
}

func (env JNIEnv) IsInstanceOf(obj JObject, clazz JClass) bool {
	return GoBool(IsInstanceOf(env.Peer(), obj.Peer(), clazz.Peer()))
}

// jni.h:
//     jmethodID (JNICALL *GetMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);

// jni.h:
//     jobject (JNICALL *CallObjectMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jobject (JNICALL *CallObjectMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallObjectMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

// jni.h:
//     jboolean (JNICALL *CallBooleanMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jboolean (JNICALL *CallBooleanMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallBooleanMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

// jni.h:
//     jbyte (JNICALL *CallByteMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jbyte (JNICALL *CallByteMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallByteMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

// jni.h:
//     jchar (JNICALL *CallCharMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jchar (JNICALL *CallCharMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallCharMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

// jni.h:
//     jshort (JNICALL *CallShortMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jshort (JNICALL *CallShortMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallShortMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

// jni.h:
//     jint (JNICALL *CallIntMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jint (JNICALL *CallIntMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jint (JNICALL *CallIntMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

// jni.h:
//     jlong (JNICALL *CallLongMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jlong (JNICALL *CallLongMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallLongMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

// jni.h:
//     jfloat (JNICALL *CallFloatMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jfloat (JNICALL *CallFloatMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallFloatMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

// jni.h:
//     jdouble (JNICALL *CallDoubleMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jdouble (JNICALL *CallDoubleMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallDoubleMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);

// jni.h:
//     void (JNICALL *CallVoidMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     void (JNICALL *CallVoidMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     void (JNICALL *CallVoidMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);

// jni.h:
//     jobject (JNICALL *CallNonvirtualObjectMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *CallNonvirtualObjectMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallNonvirtualObjectMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);

// jni.h:
//     jboolean (JNICALL *CallNonvirtualBooleanMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jboolean (JNICALL *CallNonvirtualBooleanMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallNonvirtualBooleanMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);

// jni.h:
//     jbyte (JNICALL *CallNonvirtualByteMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jbyte (JNICALL *CallNonvirtualByteMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallNonvirtualByteMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jchar (JNICALL *CallNonvirtualCharMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jchar (JNICALL *CallNonvirtualCharMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallNonvirtualCharMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jshort (JNICALL *CallNonvirtualShortMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jshort (JNICALL *CallNonvirtualShortMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallNonvirtualShortMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jint (JNICALL *CallNonvirtualIntMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jint (JNICALL *CallNonvirtualIntMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jint (JNICALL *CallNonvirtualIntMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jlong (JNICALL *CallNonvirtualLongMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jlong (JNICALL *CallNonvirtualLongMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallNonvirtualLongMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jfloat (JNICALL *CallNonvirtualFloatMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jfloat (JNICALL *CallNonvirtualFloatMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallNonvirtualFloatMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jdouble (JNICALL *CallNonvirtualDoubleMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jdouble (JNICALL *CallNonvirtualDoubleMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallNonvirtualDoubleMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     void (JNICALL *CallNonvirtualVoidMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     void (JNICALL *CallNonvirtualVoidMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     void (JNICALL *CallNonvirtualVoidMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);

// jni.h:
//     jfieldID (JNICALL *GetFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);

// jni.h:
//     jobject (JNICALL *GetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jboolean (JNICALL *GetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jbyte (JNICALL *GetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jchar (JNICALL *GetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jshort (JNICALL *GetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jint (JNICALL *GetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jlong (JNICALL *GetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jfloat (JNICALL *GetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID);
// jni.h:
//     jdouble (JNICALL *GetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID);

// jni.h:
//     void (JNICALL *SetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
// jni.h:
//     void (JNICALL *SetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
// jni.h:
//     void (JNICALL *SetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
// jni.h:
//     void (JNICALL *SetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
// jni.h:
//     void (JNICALL *SetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
// jni.h:
//     void (JNICALL *SetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
// jni.h:
//     void (JNICALL *SetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
// jni.h:
//     void (JNICALL *SetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
// jni.h:
//     void (JNICALL *SetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);

// jni.h:
//     jmethodID (JNICALL *GetStaticMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);

// jni.h:
//     jobject (JNICALL *CallStaticObjectMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *CallStaticObjectMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallStaticObjectMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jboolean (JNICALL *CallStaticBooleanMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jboolean (JNICALL *CallStaticBooleanMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallStaticBooleanMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jbyte (JNICALL *CallStaticByteMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jbyte (JNICALL *CallStaticByteMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallStaticByteMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jchar (JNICALL *CallStaticCharMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jchar (JNICALL *CallStaticCharMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallStaticCharMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jshort (JNICALL *CallStaticShortMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jshort (JNICALL *CallStaticShortMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallStaticShortMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jint (JNICALL *CallStaticIntMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jint (JNICALL *CallStaticIntMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jint (JNICALL *CallStaticIntMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jlong (JNICALL *CallStaticLongMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jlong (JNICALL *CallStaticLongMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallStaticLongMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jfloat (JNICALL *CallStaticFloatMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jfloat (JNICALL *CallStaticFloatMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallStaticFloatMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jdouble (JNICALL *CallStaticDoubleMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jdouble (JNICALL *CallStaticDoubleMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallStaticDoubleMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     void (JNICALL *CallStaticVoidMethod)(JNIEnv *env, jclass cls, jmethodID methodID, ...);
//     void (JNICALL *CallStaticVoidMethodV)(JNIEnv *env, jclass cls, jmethodID methodID, va_list args);
//     void (JNICALL *CallStaticVoidMethodA)(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);

// jni.h:
//     jfieldID (JNICALL *GetStaticFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
// jni.h:
//     jobject (JNICALL *GetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jboolean (JNICALL *GetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jbyte (JNICALL *GetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jchar (JNICALL *GetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jshort (JNICALL *GetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jint (JNICALL *GetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jlong (JNICALL *GetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jfloat (JNICALL *GetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jni.h:
//     jdouble (JNICALL *GetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID);

// jni.h:
//     void (JNICALL *SetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
// jni.h:
//     void (JNICALL *SetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
// jni.h:
//     void (JNICALL *SetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
// jni.h:
//     void (JNICALL *SetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
// jni.h:
//     void (JNICALL *SetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
// jni.h:
//     void (JNICALL *SetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
// jni.h:
//     void (JNICALL *SetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
// jni.h:
//     void (JNICALL *SetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
// jni.h:
//     void (JNICALL *SetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);

// jni.h:
//     jstring (JNICALL *NewString)(JNIEnv *env, const jchar *unicode, jsize len);
// jni.h:
//     jsize (JNICALL *GetStringLength)(JNIEnv *env, jstring str);
// jni.h:
//     const jchar *(JNICALL *GetStringChars)(JNIEnv *env, jstring str, jboolean *isCopy);
// jni.h:
//     void (JNICALL *ReleaseStringChars)(JNIEnv *env, jstring str, const jchar *chars);

// jni.h:
//     jstring (JNICALL *NewStringUTF)(JNIEnv *env, const char *utf);
// jni.h:
//     jsize (JNICALL *GetStringUTFLength)(JNIEnv *env, jstring str);
// jni.h:
//     const char* (JNICALL *GetStringUTFChars)(JNIEnv *env, jstring str, jboolean *isCopy);
// jni.h:
//     void (JNICALL *ReleaseStringUTFChars)(JNIEnv *env, jstring str, const char* chars);

// jni.h:
//     jsize (JNICALL *GetArrayLength)(JNIEnv *env, jarray array);

// jni.h:
//     jobjectArray (JNICALL *NewObjectArray)(JNIEnv *env, jsize len, jclass clazz, jobject init);
// jni.h:
//     jobject (JNICALL *GetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index);
// jni.h:
//     void (JNICALL *SetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index, jobject val);

// jni.h:
//     jbooleanArray (JNICALL *NewBooleanArray)(JNIEnv *env, jsize len);
// jni.h:
//     jbyteArray (JNICALL *NewByteArray)(JNIEnv *env, jsize len);
// jni.h:
//     jcharArray (JNICALL *NewCharArray)(JNIEnv *env, jsize len);
// jni.h:
//     jshortArray (JNICALL *NewShortArray)(JNIEnv *env, jsize len);
// jni.h:
//     jintArray (JNICALL *NewIntArray)(JNIEnv *env, jsize len);
// jni.h:
//     jlongArray (JNICALL *NewLongArray)(JNIEnv *env, jsize len);
// jni.h:
//     jfloatArray (JNICALL *NewFloatArray)(JNIEnv *env, jsize len);
// jni.h:
//     jdoubleArray (JNICALL *NewDoubleArray)(JNIEnv *env, jsize len);

// jni.h:
//     jboolean * (JNICALL *GetBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *isCopy);
// jni.h:
//     jbyte * (JNICALL *GetByteArrayElements)(JNIEnv *env, jbyteArray array, jboolean *isCopy);
// jni.h:
//     jchar * (JNICALL *GetCharArrayElements)(JNIEnv *env, jcharArray array, jboolean *isCopy);
// jni.h:
//     jshort * (JNICALL *GetShortArrayElements)(JNIEnv *env, jshortArray array, jboolean *isCopy);
// jni.h:
//     jint * (JNICALL *GetIntArrayElements)(JNIEnv *env, jintArray array, jboolean *isCopy);
// jni.h:
//     jlong * (JNICALL *GetLongArrayElements)(JNIEnv *env, jlongArray array, jboolean *isCopy);
// jni.h:
//     jfloat * (JNICALL *GetFloatArrayElements)(JNIEnv *env, jfloatArray array, jboolean *isCopy);
// jni.h:
//     jdouble * (JNICALL *GetDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jboolean *isCopy);

// jni.h:
//     void (JNICALL *ReleaseBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
// jni.h:
//     void (JNICALL *ReleaseByteArrayElements)(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
// jni.h:
//     void (JNICALL *ReleaseCharArrayElements)(JNIEnv *env, jcharArray array, jchar *elems, jint mode);
// jni.h:
//     void (JNICALL *ReleaseShortArrayElements)(JNIEnv *env, jshortArray array, jshort *elems, jint mode);
// jni.h:
//     void (JNICALL *ReleaseIntArrayElements)(JNIEnv *env, jintArray array, jint *elems, jint mode);
// jni.h:
//     void (JNICALL *ReleaseLongArrayElements)(JNIEnv *env, jlongArray array, jlong *elems, jint mode);
// jni.h:
//     void (JNICALL *ReleaseFloatArrayElements)(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
// jni.h:
//     void (JNICALL *ReleaseDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);

// jni.h:
//     void (JNICALL *GetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
// jni.h:
//     void (JNICALL *GetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
// jni.h:
//     void (JNICALL *GetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
// jni.h:
//     void (JNICALL *GetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
// jni.h:
//     void (JNICALL *GetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
// jni.h:
//     void (JNICALL *GetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
// jni.h:
//     void (JNICALL *GetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
// jni.h:
//     void (JNICALL *GetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);

// jni.h:
//     void (JNICALL *SetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);
// jni.h:
//     void (JNICALL *SetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);
// jni.h:
//     void (JNICALL *SetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);
// jni.h:
//     void (JNICALL *SetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);
// jni.h:
//     void (JNICALL *SetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
// jni.h:
//     void (JNICALL *SetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);
// jni.h:
//     void (JNICALL *SetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);
// jni.h:
//     void (JNICALL *SetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);

// jni.h:
//     jint (JNICALL *RegisterNatives)(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods);
// jni.h:
//     jint (JNICALL *UnregisterNatives)(JNIEnv *env, jclass clazz);

// jni.h:
//     jint (JNICALL *MonitorEnter)(JNIEnv *env, jobject obj);
// jni.h:
//     jint (JNICALL *MonitorExit)(JNIEnv *env, jobject obj);

// jni.h:
//     jint (JNICALL *GetJavaVM)(JNIEnv *env, JavaVM **vm);

// jni.h:
//     void (JNICALL *GetStringRegion)(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
// jni.h:
//     void (JNICALL *GetStringUTFRegion)(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);

// jni.h:
//     void * (JNICALL *GetPrimitiveArrayCritical)(JNIEnv *env, jarray array, jboolean *isCopy);
// jni.h:
//     void (JNICALL *ReleasePrimitiveArrayCritical)(JNIEnv *env, jarray array, void *carray, jint mode);

// jni.h:
//     const jchar * (JNICALL *GetStringCritical)(JNIEnv *env, jstring string, jboolean *isCopy);
// jni.h:
//     void (JNICALL *ReleaseStringCritical)(JNIEnv *env, jstring string, const jchar *cstring);

// jni.h:
//     jweak (JNICALL *NewWeakGlobalRef)(JNIEnv *env, jobject obj);
// jni.h:
//     void (JNICALL *DeleteWeakGlobalRef)(JNIEnv *env, jweak ref);

// jni.h:
//     jboolean (JNICALL *ExceptionCheck)(JNIEnv *env);

// jni.h:
//     jobject (JNICALL *NewDirectByteBuffer)(JNIEnv* env, void* address, jlong capacity);
// jni.h:
//     void* (JNICALL *GetDirectBufferAddress)(JNIEnv* env, jobject buf);
// jni.h:
//     jlong (JNICALL *GetDirectBufferCapacity)(JNIEnv* env, jobject buf);

//     /* New JNI 1.6 Features */

// jni.h:
//     jobjectRefType (JNICALL *GetObjectRefType)(JNIEnv* env, jobject obj);

// jni.h:
//     typedef struct JavaVMOption {
//         char *optionString;
//         void *extraInfo;
//     } JavaVMOption;

// jni.h:
//     typedef struct JavaVMInitArgs {
//         jint version;
//
//         jint nOptions;
//         JavaVMOption *options;
//         jboolean ignoreUnrecognized;
//     } JavaVMInitArgs;

// jni.h:
//     typedef struct JavaVMAttachArgs {
//         jint version;
//
//         char *name;
//         jobject group;
//     } JavaVMAttachArgs;

// jni.h:
//     jint (JNICALL *DestroyJavaVM)(JavaVM *vm);

// jni.h:
//     jint (JNICALL *AttachCurrentThread)(JavaVM *vm, void **penv, void *args);

// jni.h:
//     jint (JNICALL *DetachCurrentThread)(JavaVM *vm);

// jni.h:
//     jint (JNICALL *GetEnv)(JavaVM *vm, void **penv, jint version);

// jni.h:
//     jint (JNICALL *AttachCurrentThreadAsDaemon)(JavaVM *vm, void **penv, void *args);

// jni.h:
//     _JNI_IMPORT_OR_EXPORT_ jint JNICALL JNI_GetDefaultJavaVMInitArgs(void *args);
// jni.h:
//     _JNI_IMPORT_OR_EXPORT_ jint JNICALL JNI_CreateJavaVM(JavaVM **pvm, void **penv, void *args);
// jni.h:
//     _JNI_IMPORT_OR_EXPORT_ jint JNICALL JNI_GetCreatedJavaVMs(JavaVM **, jsize, jsize *);

/* Defined by native libraries. */
// jni.h:
//     JNIEXPORT jint JNICALL JNI_OnLoad(JavaVM *vm, void *reserved);
// jni.h:
//     JNIEXPORT void JNICALL JNI_OnUnload(JavaVM *vm, void *reserved);

// jni.h:
//     #define JNI_VERSION_1_1 0x00010001
//     #define JNI_VERSION_1_2 0x00010002
//     #define JNI_VERSION_1_4 0x00010004
//     #define JNI_VERSION_1_6 0x00010006
//     #define JNI_VERSION_1_8 0x00010008
const (
	JNI_VERSION_1_1 = C.JNI_VERSION_1_1
	JNI_VERSION_1_2 = C.JNI_VERSION_1_2
	JNI_VERSION_1_4 = C.JNI_VERSION_1_4
	JNI_VERSION_1_6 = C.JNI_VERSION_1_6
	JNI_VERSION_1_8 = C.JNI_VERSION_1_8
)

/* Support Functions */

func JavaBoolean(b bool) C.jboolean {
	if b {
		return JNI_TRUE
	}
	return JNI_FALSE
}

func GoBool(z C.jboolean) bool {
	return z != JNI_FALSE
}

// For example, see:
//     src/net/cgo_unix.go
//     src/os/user/lookup_unix.go
func WithCString(s string, f func(*C.char)) {
	pc := C.CString(s)
	defer C.free(unsafe.Pointer(pc))
	f(pc)
}