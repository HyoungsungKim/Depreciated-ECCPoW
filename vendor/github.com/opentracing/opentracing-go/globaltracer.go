package opentracing

<<<<<<< HEAD
var (
	globalTracer Tracer = NoopTracer{}
=======
type registeredTracer struct {
	tracer       Tracer
	isRegistered bool
}

var (
	globalTracer = registeredTracer{NoopTracer{}, false}
>>>>>>> upstream/master
)

// SetGlobalTracer sets the [singleton] opentracing.Tracer returned by
// GlobalTracer(). Those who use GlobalTracer (rather than directly manage an
// opentracing.Tracer instance) should call SetGlobalTracer as early as
// possible in main(), prior to calling the `StartSpan` global func below.
// Prior to calling `SetGlobalTracer`, any Spans started via the `StartSpan`
// (etc) globals are noops.
func SetGlobalTracer(tracer Tracer) {
<<<<<<< HEAD
	globalTracer = tracer
=======
	globalTracer = registeredTracer{tracer, true}
>>>>>>> upstream/master
}

// GlobalTracer returns the global singleton `Tracer` implementation.
// Before `SetGlobalTracer()` is called, the `GlobalTracer()` is a noop
// implementation that drops all data handed to it.
func GlobalTracer() Tracer {
<<<<<<< HEAD
	return globalTracer
=======
	return globalTracer.tracer
>>>>>>> upstream/master
}

// StartSpan defers to `Tracer.StartSpan`. See `GlobalTracer()`.
func StartSpan(operationName string, opts ...StartSpanOption) Span {
<<<<<<< HEAD
	return globalTracer.StartSpan(operationName, opts...)
=======
	return globalTracer.tracer.StartSpan(operationName, opts...)
>>>>>>> upstream/master
}

// InitGlobalTracer is deprecated. Please use SetGlobalTracer.
func InitGlobalTracer(tracer Tracer) {
	SetGlobalTracer(tracer)
}
<<<<<<< HEAD
=======

// IsGlobalTracerRegistered returns a `bool` to indicate if a tracer has been globally registered
func IsGlobalTracerRegistered() bool {
	return globalTracer.isRegistered
}
>>>>>>> upstream/master
