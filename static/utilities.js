export { getDateString, getTimeString, localTimeString };

function getDateString(dt) {
    const year = dt.getFullYear()
    const m = dt.getMonth() + 1
    var month = "" + m
    if (m < 10) {
      month = "0" + m
    }
    const d = dt.getDate()
    var day = "" + d
    if (d < 10) {
      day = "0" + d
    }
    return (day+"."+month+"."+year)
  }
  function getTimeString(dt) {
    const h = dt.getHours()
    var hour = "" + h
    if (h < 10) {
      hour = "0" + h
    }
    const m = dt.getMinutes()
    var min = "" + m
    if (m < 10) {
      min = "0" + m
    }
    return (hour+":"+min)
  } 
  function diffSeconds(dt2, dt1) {
    var diff =(dt2.getTime() - dt1.getTime()) / 1000;

    return Math.abs(Math.round(diff));
  }
  function localTimeString(dateTimeString) {
    const now = new Date();
    const then = new Date(dateTimeString)
    if (now.toLocaleDateString() == then.toLocaleDateString()) {
      const dSeconds = diffSeconds(now, then)
      if (dSeconds < 60) {
        return(dSeconds + " sec ago")
      }
      if (dSeconds < 60 * 5) {
        return(Math.abs(Math.round(dSeconds / 60)) + " min ago")
      }
      return(getTimeString(then))
    }
    return(getDateString(then) + " " + getTimeString(then))
  }