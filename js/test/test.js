var eventHandler = function(event, type) {
  if (!event.preventDefault) {
    event.preventDefault = function() {
      event.returnValue = false; //ie
    };
  }

  if (!event.stopPropagation) {
    event.stopPropagation = function() {
      event.cancelBubble = true; //ie
    };
  }

  if (!event.target) {
    event.target = event.srcElement || document;
  }

  if (isUndefined(event.defaultPrevented)) {
    var prevent = event.preventDefault;
    event.preventDefault = function() {
      event.defaultPrevented = true;
      prevent.call(event);
    };
    event.defaultPrevented = false;
  }

  event.isDefaultPrevented = function() {
    return event.defaultPrevented;
  };

  forEach(events[type || event.type], function(fn) {
    fn.call(element, event);
  });

  // Remove monkey-patched methods (IE),
  // as they would cause memory leaks in IE8.
  if (msie <= 8) {
    // IE7/8 does not allow to delete property on native object
    event.preventDefault = null;
    event.stopPropagation = null;
    event.isDefaultPrevented = null;
  } else {
    // It shouldn't affect normal browsers (native methods are defined on prototype).
    delete event.preventDefault;
    delete event.stopPropagation;
    delete event.isDefaultPrevented;
  }
}