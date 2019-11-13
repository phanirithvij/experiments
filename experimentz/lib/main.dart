// import 'dart:math';

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
  PageController _rowController;
  // No need to use valuenotifiers here
  // But I need pass by reference thus using it as a wrapper
  ValueNotifier<int> currIdxNotifier = ValueNotifier(0);
  ValueNotifier<int> currUpNotifier = ValueNotifier(0);

  @override
  void initState() {
    _controllers = [
      PageController(),
      PageController(),
      PageController(),
    ];
    _rowController = PageController(
        // initialPage: 0,
        // keepPage: true,
        );
    currIdxNotifier.value = _rowController.initialPage;
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: PageView(
        // pageSnapping: true,
        controller: _rowController,
        onPageChanged: (pno) {
          // MUST set state and trigger a rebuild
          // As horizontal viewport changed
          setState(() {
            currIdxNotifier.value = pno;
            // print("${currIdxNotifier.value} horizz");
          });
        },
        children: [
          ColPageView(
            idx: 0,
            currup: currUpNotifier,
            notifier: currIdxNotifier,
            controllers: _controllers,
            children: <Widget>[
              ColoredWidget(
                color: Colors.orange[50],
                text: "0, 0",
              ),
              ColoredWidget(
                color: Colors.orange[100],
                text: "0, 1",
              ),
              ColoredWidget(
                color: Colors.orange[200],
                text: "0, 2",
              ),
              ColoredWidget(
                color: Colors.orange[300],
                text: "0, 3",
              ),
            ],
          ),
          ColPageView(
            idx: 1,
            currup: currUpNotifier,
            notifier: currIdxNotifier,
            controllers: _controllers,
            children: [
              ColoredWidget(
                color: Colors.green[100],
                text: "1, 0",
              ),
            ],
          ),
          ColPageView(
            idx: 2,
            currup: currUpNotifier,
            notifier: currIdxNotifier,
            controllers: _controllers,
            children: [
              ColoredWidget(
                color: Colors.teal[100],
                text: "2, 0",
              ),
              ColoredWidget(
                color: Colors.teal[200],
                text: "2, 1",
              ),
              ColoredWidget(
                color: Colors.teal[300],
                text: "2, 2",
              ),
              ColoredWidget(
                color: Colors.teal[400],
                text: "2, 3",
              ),
            ],
          ),
        ],
      ),
    );
  }
}

/// All the vertical pageviews are here
class ColPageView extends StatefulWidget {
  final List<Widget> children;
  final List<PageController> controllers;
  final ValueNotifier<int> notifier;
  final ValueNotifier<int> currup;
  final int idx;

  const ColPageView({
    Key key,
    this.children = const <Widget>[],
    @required this.idx,
    @required this.currup,
    @required this.notifier,
    @required this.controllers,
  }) : super(key: key);

  @override
  _ColPageViewState createState() => _ColPageViewState();
}

class _ColPageViewState extends State<ColPageView> {
  @override
  void initState() {
    // Just initialized
    // Set the start value to be the current vertical value
    widget.controllers[widget.idx] = PageController(
      initialPage: widget.currup.value ?? 0,
      keepPage: true,
    );
    // print("INIT STATE ${widget.idx}");
    // print("INIT STATE ${widget.currup.value}");
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return PageView(
      // pageSnapping: true,
      controller: widget.controllers[widget.idx],
      //   controller: widget.controller,
      scrollDirection: Axis.vertical,
      children: widget.children,
      onPageChanged: (widget.notifier.value == widget.idx)
          ? (pno) {
              // if the global horizontal page is the current widget
              // var rand = Random();
              // var randnn = rand.nextDouble();
              widget.controllers.forEach((colpv) {
                if (widget.controllers[widget.idx] == colpv) {
                  // print("same widget so return $randnn");
                  return;
                }
                // https://github.com/flutter/flutter/issues/20621#issuecomment-445504085
                // Only if the controller has clients
                bool isSelected = colpv.hasClients
                    ? colpv.page == pno
                    : colpv.initialPage == pno;

                // Not the same page as everyone
                if (!isSelected) {
                  // print("not selected");
                  if (colpv.hasClients) {
                    colpv.animateToPage(
                      pno,
                      duration: Duration(milliseconds: 200),
                      curve: Curves.easeIn,
                    );
                  }
                }
                // set the current updated value of the vertical coord
                widget.currup.value = pno;
                // print("$pno $isSelected");
              });
              // print("col-${widget.idx} changed to $pno");

              // set horizontal coord to be null
              // As we've finished dealing with it
              widget.notifier.value = null;
            }
          : (_) {
              // Others which are not the currently moving pageview
              // SHOULD not have any listeners
              // Spent 5hrs trying to figure this out
              // print("nope ${widget.notifier.value} == ${widget.idx}");
            },
    );
  }
}

/// A Widget that simply displays a color and an input text
/// NOTE: This is a StatefulWidget because needs to use keepalive
class ColoredWidget extends StatefulWidget {
  /// Color to display the widget in
  final Color color;
  final String text;

  const ColoredWidget({
    Key key,
    @required this.color,
    @required this.text,
  }) : super(key: key);

  @override
  _ColoredWidgetState createState() => _ColoredWidgetState();
}

class _ColoredWidgetState extends State<ColoredWidget>
    with AutomaticKeepAliveClientMixin<ColoredWidget> {
  @override
  Widget build(BuildContext context) {
    super.build(context);
    return Container(
        color: widget.color,
        child: Center(
          child: Text(
            widget.text,
            style: TextStyle(
              fontSize: 100,
              color: Colors.black,
            ),
          ),
        ));
  }

  // Need to use this or the state of the pageview will be lost
  // In this case if not using keepalive it would still work but
  // It will scroll down everytime the page gets changed horizontally
  // TODO: Destroy if not next to current pageview
  @override
  bool get wantKeepAlive => true;
}
