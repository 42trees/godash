(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['text'] = template({"compiler":[6,">= 2.0.0-beta.1"],"main":function(depth0,helpers,partials,data) {
  var helper, functionType="function", helperMissing=helpers.helperMissing, escapeExpression=this.escapeExpression;
  return "\n<div class=\"widget pure-u-1 pure-u-md-1-3\" id=\"widget-"
    + escapeExpression(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : helperMissing),(typeof helper === functionType ? helper.call(depth0, {"name":"Id","hash":{},"data":data}) : helper)))
    + "\">\n  <div class=\"widget-contents "
    + escapeExpression(((helper = (helper = helpers.Color || (depth0 != null ? depth0.Color : depth0)) != null ? helper : helperMissing),(typeof helper === functionType ? helper.call(depth0, {"name":"Color","hash":{},"data":data}) : helper)))
    + "\">\n    <h2>"
    + escapeExpression(((helper = (helper = helpers.Title || (depth0 != null ? depth0.Title : depth0)) != null ? helper : helperMissing),(typeof helper === functionType ? helper.call(depth0, {"name":"Title","hash":{},"data":data}) : helper)))
    + " </h2>\n    <h3>"
    + escapeExpression(((helper = (helper = helpers.Subtitle || (depth0 != null ? depth0.Subtitle : depth0)) != null ? helper : helperMissing),(typeof helper === functionType ? helper.call(depth0, {"name":"Subtitle","hash":{},"data":data}) : helper)))
    + "</h3>  \n    <p>"
    + escapeExpression(((helper = (helper = helpers.Content || (depth0 != null ? depth0.Content : depth0)) != null ? helper : helperMissing),(typeof helper === functionType ? helper.call(depth0, {"name":"Content","hash":{},"data":data}) : helper)))
    + "</p>\n    <h5>"
    + escapeExpression(((helper = (helper = helpers.Time || (depth0 != null ? depth0.Time : depth0)) != null ? helper : helperMissing),(typeof helper === functionType ? helper.call(depth0, {"name":"Time","hash":{},"data":data}) : helper)))
    + " </h5>\n  </div>\n</div>\n\n";
},"useData":true});
})();