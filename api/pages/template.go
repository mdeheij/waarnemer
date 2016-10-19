package pages

import "html/template"

//Template contains a HTML page with a logo in SVG format
var Template = template.Must(template.New("routes").Parse(`
<html>
<head>
<title>Monitoring API</title>
<style>
    html, body {
        background-color: #333;
        margin: 0;
        padding: 0;
        font-family: sans-serif;
    }
    svg {
        position: fixed;
        top: 25%;
        left: 25%;
    }
</style>
</head>
<body>
<svg
   xmlns:dc="http://purl.org/dc/elements/1.1/"
   xmlns:cc="http://creativecommons.org/ns#"
   xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
   xmlns:svg="http://www.w3.org/2000/svg"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
   viewBox="0 0 48 48"
   id="svg3072"
   version="1.1"
   inkscape:version="0.48.4 r9939"
   width="50%"
   height="50%"
   sodipodi:docname="monitoring.svg">
  <metadata
     id="metadata3119">
    <rdf:RDF>
      <cc:Work
         rdf:about="">
        <dc:format>image/svg+xml</dc:format>
        <dc:type
           rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
        <dc:title></dc:title>
      </cc:Work>
    </rdf:RDF>
  </metadata>
  <sodipodi:namedview
     pagecolor="#ffffff"
     bordercolor="#666666"
     borderopacity="1"
     objecttolerance="10"
     gridtolerance="10"
     guidetolerance="10"
     inkscape:pageopacity="0"
     inkscape:pageshadow="2"
     inkscape:window-width="1920"
     inkscape:window-height="1056"
     id="namedview3117"
     showgrid="false"
     inkscape:snap-global="true"
     inkscape:zoom="9.46875"
     inkscape:cx="63.095117"
     inkscape:cy="15.256255"
     inkscape:window-x="1920"
     inkscape:window-y="24"
     inkscape:window-maximized="1"
     inkscape:current-layer="g3531" />
  <defs
     id="defs3074">
    <linearGradient
       id="linearGradient3764"
       x1="1"
       x2="47"
       gradientUnits="userSpaceOnUse"
       gradientTransform="matrix(0,-1,1,0,-1.5e-6,47.999998)">
      <stop
         style="stop-color:#2290ee;stop-opacity:1"
         id="stop3077" />
      <stop
         offset="1"
         style="stop-color:#359bef;stop-opacity:1"
         id="stop3079" />
    </linearGradient>
    <filter
       color-interpolation-filters="sRGB"
       id="color-overlay-1"
       filterUnits="userSpaceOnUse">
      <feFlood
         flood-color="#868889"
         id="feFlood3217" />
      <feComposite
         operator="in"
         in2="SourceGraphic"
         id="feComposite3219" />
      <feBlend
         mode="normal"
         in2="SourceGraphic"
         result="solidFill"
         id="feBlend3221" />
    </filter>
    <filter
       color-interpolation-filters="sRGB"
       id="color-overlay-1-5"
       filterUnits="userSpaceOnUse">
      <feFlood
         flood-color="#868889"
         id="feFlood3217-6" />
      <feComposite
         operator="in"
         in2="SourceGraphic"
         id="feComposite3219-2" />
      <feBlend
         mode="normal"
         in2="SourceGraphic"
         result="solidFill"
         id="feBlend3221-9" />
    </filter>
  </defs>
  <g
     id="g3081">
    <path
       d="m 36.31 5 c 5.859 4.062 9.688 10.831 9.688 18.5 c 0 12.426 -10.07 22.5 -22.5 22.5 c -7.669 0 -14.438 -3.828 -18.5 -9.688 c 1.037 1.822 2.306 3.499 3.781 4.969 c 4.085 3.712 9.514 5.969 15.469 5.969 c 12.703 0 23 -10.298 23 -23 c 0 -5.954 -2.256 -11.384 -5.969 -15.469 c -1.469 -1.475 -3.147 -2.744 -4.969 -3.781 z m 4.969 3.781 c 3.854 4.113 6.219 9.637 6.219 15.719 c 0 12.703 -10.297 23 -23 23 c -6.081 0 -11.606 -2.364 -15.719 -6.219 c 4.16 4.144 9.883 6.719 16.219 6.719 c 12.703 0 23 -10.298 23 -23 c 0 -6.335 -2.575 -12.06 -6.719 -16.219 z"
       style="opacity:0.05"
       id="path3083" />
    <path
       d="m 41.28 8.781 c 3.712 4.085 5.969 9.514 5.969 15.469 c 0 12.703 -10.297 23 -23 23 c -5.954 0 -11.384 -2.256 -15.469 -5.969 c 4.113 3.854 9.637 6.219 15.719 6.219 c 12.703 0 23 -10.298 23 -23 c 0 -6.081 -2.364 -11.606 -6.219 -15.719 z"
       style="opacity:0.1"
       id="path3085" />
    <path
       d="m 31.25 2.375 c 8.615 3.154 14.75 11.417 14.75 21.13 c 0 12.426 -10.07 22.5 -22.5 22.5 c -9.708 0 -17.971 -6.135 -21.12 -14.75 a 23 23 0 0 0 44.875 -7 a 23 23 0 0 0 -16 -21.875 z"
       style="opacity:0.2"
       id="path3087" />
  </g>
  <g
     id="g3089">
    <path
       d="m 24 1 c 12.703 0 23 10.297 23 23 c 0 12.703 -10.297 23 -23 23 -12.703 0 -23 -10.297 -23 -23 0 -12.703 10.297 -23 23 -23 z"
       style="fill:url(#linearGradient3764);fill-opacity:1"
       id="path3091" />
  </g>
  <g
     id="g3113">
    <path
       d="m 40.03 7.531 c 3.712 4.084 5.969 9.514 5.969 15.469 0 12.703 -10.297 23 -23 23 c -5.954 0 -11.384 -2.256 -15.469 -5.969 4.178 4.291 10.01 6.969 16.469 6.969 c 12.703 0 23 -10.298 23 -23 0 -6.462 -2.677 -12.291 -6.969 -16.469 z"
       style="opacity:0.1"
       id="path3115" />
  </g>
  <g
     id="g3531"
     transform="matrix(0.45274528,0,0,0.45274528,8.4288342,-6.2110734)">
    <polyline
       id="polyline3407"
       points="  32.3,77.6 32.3,64.8 95.8,64.8 95.8,77.6 "
       stroke-miterlimit="22.9256"
       style="fill:none;stroke:#414b5e;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:22.92560005"
       transform="translate(-29.78218,4.8640249)" />
    <line
       id="line3409"
       y2="60.064026"
       x2="34.217819"
       y1="81.76403"
       x1="34.217819"
       stroke-miterlimit="22.9256"
       style="fill:none;stroke:#414b5e;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:22.92560005" />
    <circle
       id="circle3411"
       r="5.1999998"
       cy="82.900002"
       cx="32.200001"
       stroke-miterlimit="22.9256"
       sodipodi:cx="32.200001"
       sodipodi:cy="82.900002"
       sodipodi:rx="5.1999998"
       sodipodi:ry="5.1999998"
       style="fill:#ffffff;stroke:#414b5e;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:22.92560004999999990"
       transform="translate(-29.78218,4.8640249)"
       d="m 37.400001,82.900002 c 0,2.87188 -2.32812,5.199999 -5.2,5.199999 -2.871881,0 -5.2,-2.328119 -5.2,-5.199999 0,-2.871881 2.328119,-5.2 5.2,-5.2 2.87188,0 5.2,2.328119 5.2,5.2 z" />
    <path
       id="path3413"
       d="m 39.31782,87.764025 c 0,2.8 -2.3,5.2 -5.2,5.2 -2.9,0 -5.2,-2.3 -5.2,-5.2 0,-2.9 2.3,-5.2 5.2,-5.2 2.9,0 5.2,2.3 5.2,5.2 z"
       stroke-miterlimit="22.9256"
       inkscape:connector-curvature="0"
       style="fill:#ffffff;stroke:#414b5e;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:22.92560004999999990" />
    <circle
       id="circle3415"
       r="5.1999998"
       cy="82.900002"
       cx="95.800003"
       stroke-miterlimit="22.9256"
       sodipodi:cx="95.800003"
       sodipodi:cy="82.900002"
       sodipodi:rx="5.1999998"
       sodipodi:ry="5.1999998"
       style="fill:#ffffff;stroke:#414b5e;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:22.92560004999999990"
       transform="translate(-29.78218,4.8640249)"
       d="m 101,82.900002 c 0,2.87188 -2.328116,5.199999 -5.199997,5.199999 -2.871881,0 -5.2,-2.328119 -5.2,-5.199999 0,-2.871881 2.328119,-5.2 5.2,-5.2 2.871881,0 5.199997,2.328119 5.199997,5.2 z" />
    <circle
       id="circle3417"
       r="12.2"
       cy="42.900002"
       cx="64"
       stroke-miterlimit="22.9256"
       sodipodi:cx="64"
       sodipodi:cy="42.900002"
       sodipodi:rx="12.2"
       sodipodi:ry="12.2"
       style="fill:#ffffff;stroke:#414b5e;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:22.92560004999999990"
       transform="translate(-29.78218,4.8640249)"
       d="m 76.2,42.900002 c 0,6.737873 -5.462126,12.199999 -12.2,12.199999 -6.737874,0 -12.2,-5.462126 -12.2,-12.199999 0,-6.737874 5.462126,-12.2 12.2,-12.2 6.737874,0 12.2,5.462126 12.2,12.2 z" />
  </g>
</svg>

</body>
</html>
`))
