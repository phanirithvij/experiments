import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

void main() {
  SystemChrome.setEnabledSystemUIOverlays([]);
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Experiments',
      theme: ThemeData.dark(),
      home: MyHomePage(title: 'FlutterExps'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  List<PageController> _controllers;
  ValueNotifier<int> _horizPage = ValueNotifier(0);
  ValueNotifier<double> _vertPage = ValueNotifier(0);

  @override
  void initState() {
    _controllers = [
      PageController(keepPage: true),
      PageController(keepPage: true),
    ];
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          ColPageView(
            idx: 0,
            currHoriz: _horizPage,
            vertstate: _vertPage,
            controllers: _controllers,
            children: <Widget>[
              ColoredWidget(
                color: Colors.cyan,
                direction: ">",
              ),
              ColoredWidget(
                color: Colors.orange,
                direction: ">>",
              ),
            ],
          ),
          ColPageView(
            idx: 1,
            currHoriz: _horizPage,
            vertstate: _vertPage,
            controllers: _controllers,
            children: [
              ColoredWidget(
                color: Colors.green,
                direction: "<",
              ),
              ColoredWidget(
                color: Colors.yellow,
                direction: "<<",
              ),
            ],
          ),
        ],
      ),
    );
  }
}

class ColPageView extends StatefulWidget {
  final int idx;
  final List<Widget> children;
  final List<PageController> controllers;
  final ValueNotifier<int> currHoriz;
  final ValueNotifier<double> vertstate;

  const ColPageView({
    Key key,
    this.children = const <Widget>[],
    @required this.controllers,
    @required this.currHoriz,
    @required this.vertstate,
    @required this.idx,
  }) : super(key: key);

  @override
  _ColPageViewState createState() => _ColPageViewState();
}

class _ColPageViewState extends State<ColPageView> {
  @override
  void initState() {
    widget.controllers[widget.idx] = PageController(
      initialPage: widget.vertstate.value.floor() ?? 0,
      keepPage: true,
    )..addListener(_onPageChanged);

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    var container = Container(
      width: MediaQuery.of(context).size.width / 2,
      child: PageView(
        controller: widget.controllers[widget.idx],
        scrollDirection: Axis.vertical,
        children: widget.children,
      ),
    );
    return container;
  }

  void _onPageChanged() {
    // print("Curr horiz ${widget.currHoriz.value} ${widget.vertstate.value}");
    // print(
    //     "onPageCallback ${widget.controllers[widget.idx].page} for ${widget.idx}");
    // if (widget.currHoriz.value == null) {
    //   widget.currHoriz.value = widget.idx;
    // }
    if (widget.idx == widget.currHoriz.value) {
      widget.controllers.forEach((colpv) {
        if (colpv != widget.controllers[widget.idx]) {
          // print("has clients and ${colpv.hasClients}");
          if (colpv.hasClients) {
            // print(colpv);
            // colpv.position.setPixels(widget.controllers[widget.idx].page);
            colpv.jumpTo(
              widget.controllers[widget.idx].page,
              // curve: Curves.easeIn,
              // duration: Duration(milliseconds: 300),
            );
            // setState(() {});
          }
        }
      });
      // set horizontal coord to be this
      // widget.currHoriz.value = null;
      // Set latest vertical position
      widget.vertstate.value = widget.controllers[widget.idx].page;

      // print("col-${widget.idx} changed to $pno");
    } else {
      // print("other widget");
      // print("${widget.controllers[widget.idx].page}");
      return null;
    }
  }
}

class ColoredWidget extends StatefulWidget {
  final Color color;
  final String direction;

  const ColoredWidget({
    Key key,
    @required this.color,
    @required this.direction,
  }) : super(key: key);

  @override
  _ColoredWidgetState createState() => _ColoredWidgetState();
}

class _ColoredWidgetState extends State<ColoredWidget> {
  @override
  Widget build(BuildContext context) {
    return Container(
        color: widget.color,
        child: Center(
          child: Text(
            widget.direction,
            style: TextStyle(
              fontSize: 100,
              color: Colors.black,
            ),
          ),
        ));
  }
}
