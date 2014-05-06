Sparkline = function(opts) {
    this.node = opts.node;
    this.data = [];
    this.numPoints = opts.numPoints || 30;
    this.sparkOpts = opts.sparkOpts;
};

Sparkline.prototype.addPoint = function(pt) {
    this.data.push(pt);
    if (this.data.length > this.numPoints) {
       this.data.splice(0, 1);
    }

    // redraw
    this.node.sparkline(this.data, this.sparkOpts);
};

jQuery(window).bind("unload", function() {
    jQuery("*").add(document).unbind();
});

setTimeout(function() {
    window.location.reload();
}, 21600000);

(function ($) {
    var graphs = {};

    $('.service-spark').each(function(i, el) {
        var slug = $(this).data('slug');
        graphs[slug] = new Sparkline({
            node: $(this)
          , numPoints: 60
          , sparkOpts: {
              lineColor: "#1c1c88"
            , fillColor: false
            , width: "300px"
          }
        });
    });

    var S = WipesClient({
      log: (console && console.log) ? console.log : function (msg) {}
      , error: (console && console.error) ? console.error : function (msg) {}
      , handleJson: function (json) {
          var selector = ''
          , service = json['service']
          , value = json['value']
          , time = json['time'];

          if (graphs[service]) {
              graphs[service].addPoint(json['value']);
              $('#' + service.replace('graph', 'value')).text(value);
          }
      }});
})(jQuery);
