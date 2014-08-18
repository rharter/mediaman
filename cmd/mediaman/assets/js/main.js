function setSelectedLibrary(id) {
  var libs = document.querySelectorAll('nav.top-bar ul.left li');
  for (var i = 0; i < libs.length; i++) {
    var lib = libs[i];
    if (lib.dataset.id == id) {
      lib.className += 'active';
    } else {
      lib.className.replace(/(?:^|\s)active(?!\S)/g, '');
    }
  }
}